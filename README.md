# Concurrent Line-by-Line SHA256 Calculator

This Go program reads a text file line by line, computes the SHA256 checksum for each line concurrently using goroutines, and prints out the checksums in the original order of the lines.
(Implemented for [Cybozu's internship](https://cybozu.co.jp/company/job/recruitment/intern/kubernetes.html) pre-assignment)

## Usage
```
% go run main.go <filename>
```
Replace `<filename>` with the path to the input file you want to process.
