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
