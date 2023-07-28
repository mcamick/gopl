# gopl
Working project with exercises, projects, and notes from "The GO Programming Language" by Alan Donovan and Brian Kernighan

## Common Commands
---
Compiles code, links libraries, then runs resulting executable file without saving
```go
go run {file name}
```
\
Compiles code, links libraries, then saves the program without running
```go
go build {file name}
```
\
Provides formatting options and reviews code for potential errors. The [gofmt](https://pkg.go.dev/cmd/gofmt) docs provide additional details on usage and flags.
```go
gofmt [flags] [path ...]
```
## Testing Commands & Guidance
---
Runs unit tests with increased verbosity
```go
go test -v
```
\
Calculates coverage for your current unit tests and outputs to an output file
```go
go test -coverprofile=coverage.out
```
\
Opens a web browser and renders detailed results from coverage tests stored in coverage.out
```go
go tool cover -html=coverage.out
```
\
Runs benchmark tests against every benchmark in the file
```go
go test -bench=.
```

## Good External Resources
---
I used this documentation when figuring out how to format and write benchmark tests for exercise 1.3., I didn't really get what they were looking for with the references to later in the chapter and chapter 11. This documentation was succinct and clearly outlined how to build tests in golang.  
[Go Testing Tutorial](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)