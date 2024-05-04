package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
)

type Result struct {
	LineNumber int
	Checksum   string
}

func calculateInParallel(file *os.File) (map[int]string, int, error) {
	scanner := bufio.NewScanner(file)
	lineNum := 0
	resultChan := make(chan Result)
	var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		wg.Add(1)

		go func(data string, nowLine int) {
			defer wg.Done()

			checksum := sha256.Sum256([]byte(data))
			checksumHex := hex.EncodeToString(checksum[:])

			resultChan <- Result{
				LineNumber: nowLine,
				Checksum:   checksumHex,
			}
		}(line, lineNum)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var checksums = make(map[int]string)
	for r := range resultChan {
		checksums[r.LineNumber] = r.Checksum
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	return checksums, lineNum, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	checksums, fileLen, err := calculateInParallel(file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for i := 1; i <= fileLen; i++ {
		fmt.Println(checksums[i])
	}
}
