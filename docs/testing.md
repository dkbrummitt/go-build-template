# Testing and Tooling

"Simple code tends to be more faster code."
~ Me

## Tools

### Run

Run your code locally, without creating an executable

```sh
go run /path/to/main.go
```

### Code Format

How it works

```sh
gofmt -h
```

Output what code **should** look like

```sh
gofmt path/to/aFile.go
```

Output **difference** between your code and what code **should** look like

```sh
gofmt -d path/to/aFile.go
```

**Update** your code to what it **should** look like
Note: If you use VSCode this will get called automatically when you save your code.

```sh
gofmt -w path/to/aFile.go
```

### Imports

How it works

```sh
goimports -h
```

Output what imports **should** look like (add or remove)

```sh
goimports path/to/aFile.go
```

Output **difference** between your imports and what imports **should** look like

```sh
goimports -d path/to/aFile.go
```

**Update** your imports to what it **should** look like
Note: If you use VSCode this will get called automatically when you save your imports.

```sh
goimports -w path/to/aFile.go
```

### Build

Generate the executable binary. By default it will build per the OS you are building on.

```sh
go build
```

Build for a particular OS using GOOS (pronounced "Go OS")

```sh
GOOS=windows go build
```

Build for a particular Architecture using GOOS (pronounced "Go OS")

```sh
GOARCH=amd64 go build
```

#### Supported GOOS values

[More Info](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
**Bold** supported out of the box

- android
- **darwin**
- **dragonfly**
- **freebsd**
- **linux**
- **nacl**
- **netbsd**
- **openbsd**
- **plan9**
- **solaris**
- **windows**
- zos

#### Supported GOARCH values (64-bit)

**Bold** supported out of the box

- **amd64**
- **arm64**
- arm64be
- **ppc64**
- **ppc64le**
- **mips64**
- **mips64le**
- **s390x**
- sparc64

#### Supported GOARCH values (32-bit)

[More Info](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
**Bold** supported out of the box

- **386**
- **amd64p32**
- **arm**
- armbe
- **mips**
- **mipsle**
- mips64p32
- mips64p32le
- ppc
- s390
- sparc

### Install

Install executables on the `$GOPATH/bin`

### Get (packages/libs)

The git-host can be github.com, bitbucket, or even a private repo

```sh
go get git-host/path/to/repo
```

### Doc

See all documentation for a package

```sh
go doc
```

See documentation for a package

```sh
go doc PACKAGE_NAME
```

See documentation for a package's function/method

```sh
go doc PACKAGE_NAME FUNCTION_NAME
```

See documentation for all documentation your system (great for offline operation)

```sh
godoc -http :PORT
```

### Check code for unchecked errors

Before use, you must install the tool. It is not bundled in go by default.

```sh
go get -u github.com/kisielk/errcheck
```

Check errors on a specific package.

```sh
errcheck PACKAGE_NAME
```

Check errors on everything

```sh
errcheck ./...
```

### Find potential bugs

```sh
go vet
```

## Other Tools to look into

- go list
- go delve (debugger)

## Testing

Execute tests. Test files should end have a naming convention of
`_test.go`. All tests in the test go file should start with `TEST`

```sh
go test
```

### Benchmarking

Add a flag to go test to turn on benchmarks

```sh
go test -bench=.
go test -bench=. ./...
```

### Tracing CPU activities

NOTES:

- For full visability into pprof results install Graphviz `brew install graphviz`
- The browser port may change for each call of this. So, I wouldnt bother bookmarking.

Add a flag to go test to turn on benchmarks

cd into the directory you want to do tracing

```sh
cd pkg/version
go test -bench=. -trace trace.out
ls -FAlth
go tool trace -http :6060 trace.out
```

`go tool trace trace.out` opens trace results in browser to a random port
`go tool trace -http :6060 trace.out` opens trace results in browser to port 6060

## HTTP Benchmarking tool

[More Info](https://github.com/tsliwowicz/go-wrk)

Install the tool

```sh
go get github.com/tsliwowicz/go-wrk
```

Start your server

start the bench mark (no request body)

```sh
go-wrk -d DURATION  -c NUM_THREADS -M METHOD http://0.0.0.0:8080
```

start the bench mark (with request body)

```sh
go-wrk -d DURATION  -c NUM_THREADS -M METHOD -body "I am a simple body" http://0.0.0.0:8080
```

```sh
touch file.json
echo "{\"message\": \"I am a simple json\"}" > file.json
go-wrk -d DURATION  -c NUM_THREADS -M METHOD -body @file.json http://0.0.0.0:8080
```

## Instramentation

### HTTP Performance Profile

Add http/pprof to your imports to provide peformance details on your API
Adding this may marginally slow down your server

```go
import (
    // IMAGINATION HERE
    _ "net/http/pprof" // This import acts as a feature flag. The underscore (_) prevents go-imports from removing an un-used import.
    // MORE IMAGINATION HERE
)
```

Run your server

Do a load test

Navigate to browser to see  results (note server is still running)
[Performance Profile Results](http://0.0.0.0:8080/debug/pprof/)

Take a profile snapshot of your application. This command is interactive. NOTE: The following should be run while your application is under load. Otherwise, the results will be emoty.

```sh
go tool pprof --seconds=5 0.0.0.0:8080/debug/pprof/profile
```

Tracing

Spans tell **what** happened.
Logs tell you **why**

### CPU Profiling

```sh
go test -bench=. -cpuprofile cpu.prof
go tool pprof -http :6060 cpu.prof
```

At the top of the resulting graph is your apps entry point and your important functions
At the bottom are low level system calls.

Use the flame graph to nail down what is getting used, and how long you are spending there.

When doing performance improvements everything that is not a bottleneck can wait.

### Race Detector

Detect race conditions on concurrent calls. All race conoditions are dangerous. The sooner we find them the better

```sh
go test -race
```

## Summary

Performance Bottlenecks are areas to target.

Processing:

- Remove useless code
- Decouple functions
- Code Reviews
- Use Pprof

Network I/O:

- reduce the number of requests where possible.
- plan ahead of launch/go-live
- batch your requests
- minimize payload
- non-blocking
  - concurrency
  - Wait Groups

Disk I/O:

- do all that you can in memory
- package bufio
- package ioutil
- non-blocking
  - concurrency
  - Wait Groups

Contention (Concurrency):

- limit synchronization
- use execution tracer
- user workers (courser granularity)
  - more work per goroutine
  - workers pattern
- bufferred channles
- sharding

Malloc (asking for memory):

- allocate less often
- allocate slices
- recycle w/ sync.Pool
- pass []byte parameter
- value vs pointer
- heap ( :( )) vs stack (preferred over heap)

Heap v Stack:

If you pass a reference to a local variable around, then it may escape to the Heap. Making it more expensive.

- use `go build -gcflags=-m` to check what the compiler is doing.
- Look for:
  - "leaking param:foo"
  - "bar escapes to heap"
  - "baz does not escape"

Regex:

These are great but expensive, so avoid them when you can.

Use instead:

- strings.Contains
- strings.Split
- strings.ReplaceAll
- bytes.Contains
- bytes.Split
- bytes.ReplaceAll
- for loops

Function calls:

Function calls incurr a non-zero cost (AKA not free). And can be expensive if called millions of times.

- use `go build -gcflags=-m` to find where we can trim the fat
- Look for:
  - can inline compute
  - inlining call to compute

# Resources:

[Performance Tuning Go Applications on GCP](https://www.youtube.com/watch?v=b0o-xeEoug0)
[Go Tooling in Action](https://www.youtube.com/watch?v=uBjoTxosSys)

[Secrets of Successful Teamwork:Insights from Google](https://youtu.be/hHIikHJV9fI)
[Why Go is Successful](https://www.youtube.com/watch?v=k9Zbuuo51go)
[Things in Go I Never Use](https://youtu.be/5DVV36uqQ4E)
[Go Proverbs](https://youtu.be/PAAkCSZUG1c)
