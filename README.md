# Go app template build environment
[![Build Status](https://travis-ci.org/dkbrummitt/go-build-template.svg?branch=master)](https://travis-ci.org/dkbrummitt/go-build-template)

This is a skeleton project for a Go application, which captures the best build
techniques I have learned to date.  It uses a Makefile to drive the build (the
universal API to software projects) and a Dockerfile to build a docker image.

This has only been tested on Linux, and depends on Docker to build.

**This template requires Go 1.11 or higher**

## Customizing it

To use this, simply copy these files and make the following changes:

Makefile:
   - change `BIN` to your binary name
   - rename `cmd/myapp` to `cmd/$BIN`
   - change `REGISTRY` to the Docker registry you want to use
   - maybe change `SRC_DIRS` if you use some other layout
   - choose a strategy for `VERSION` values - git tags or manual

Dockerfile.in:
   - maybe change or remove the `USER` if you need

## Go Modules

This assumes the use of go modules (which will be the default for all Go builds
as of Go 1.13).
This does NOT assume the use of vendoring (which reasonable minds might disagree about).
If you wish to use vendoring, you will need to run `go mod vendor` to create a `vendor` directory when you
have dependencies. Otherwise,

```shell
go get ./...
```

will be used instead

## Building

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all architectures.

Run `make push` to push the container image to `REGISTRY`.  Run `make all-push`
to push the container images for all architectures.

Run `make clean` to clean up.

## Garden Tending

### Comments

Try to provide detailed comments when possible/relevant especially for public
functions/methods. The format below is not required, but the content described
below, offers things to consider.

```go
// FUNCTION/METHOD NAME descrition of what the method/function does
//
// Pre-Condition:
// - are there any actions/states that are needed before this is executed
// Post-Condition:
// - are there any states that are affected after this is executed
// Params:
// - describe params if any, as well as param validation
// Returns:
// - describe return values if any, as well as expectations
// Errors:
// - describe conditons(s) where an error would be retured
// Dev Notes:
// - Notes to other maintainers/cliets that may be helpful
```

### errcheck

Use a tool like errcheck to check for any unchecked errors in the code base.
Sometimes uncheck errors are intended. This tool will help detect it when
unintended.

```sh
go get -u github.com/kisielk/errcheck
errcheck ./...
```

### Unit Testing

#### What should have tests

Functions/Methods should have a unit test if they meet any of the following:

- Public facing
- A bug was found in the method or function (public/private). Add a to verify bug-fix to ensure its not re-introduced
- the complexity of the method/function is higher than 10. For both public and private.

Codacy has a great [article](https://www.codacy.com/blog/an-in-depth-explanation-of-code-complexity/) on code complexity.

Tools like [Sonarqube](https://www.sonarqube.org/), can help automate checking for code complexity. Added bonus, it supports a ridiculous number of languages.

#### Too Small To Test

Functions/Methods may be too small to test if the meet any of the following
criteria:

- do not have any logic branches (if, switch , loops)
- is a simple getter/setter expecially if it does not have any side-effects

#### Table Driven Tests
Consider using table driven testing when necessary. Its a great way to both
consolidate and outline test cases. It is also an EXCELLENT way to ensure
that a single test can cover multiple logic branches in your code.

```go

    //define the important permutations...
    tstCases := []struct {
		//place imagination here
	}{
        {},
        {},
    }

    for _,testCase := range tstCases{
        // verify your test case
    }
```

#### Executing tests

Add flags to the test call

`-failfast` to halt the tests at the first sign of trouble
`-race` to check for race conditions. ESPECIALLY if you are using concurrency.

NOTE: Adding -race can slow down test execution.

```go
go test -race -failfast ./...
```

### Benchmarking

Use benchmarking to measure how fast your application is performaing. The
variety of circumstances of when to/not-to write are benchmark are too vast,
zso I will only say, if you feel its needed, add it.

That said, here are a few flags that you may find usefule:

- `-benchtime` to specify how long the bench should run (OPTIONAL)
- `-benchmem`  to check memory during the bench testing (OPTIONAL)
- `-bench` specify the regex of what should be benchmarked. (REQUIRED)
- `count` how many times should the bench be executed (OPTIONAL)
- `-cpu=1,4,8` benchmark concurrent that are using concurrency (OPTIONAL)

```sh
# run for 20 seconds
go test -bench=. -benchtime=20s -count 3 ./...

# run for 20 iterations
go test -bench=. -benchtime=20x -count 3 ./...
```

Note, using benchtime with count is likely equivalent to mini stress testing your packages.

Use Benchcmp to compare results between benchmarks

```sh
go get golang.org/x/tools/cmd/benchcmp
go test -benchmem -bench=.  ./... > $(date '+%Y-%m-%dT%H:%M:%S').benchmark.txt
go test -benchmem -bench=.  ./... > $(date '+%Y-%m-%dT%H:%M:%S').benchmark.txt
benchcmp old.benchmark.txt new.benchmark.txt
```

Putting it all together

The hack directory has a shellscript that will called profile.sh that will run the benchmarks and captures memory and cpu profile data.

Usage:

```sh
hack/profile.sh
```

The results are stored in a directory that is created alled `generated` and is grouped by package names

#### Stress Testing

Memory and concurrency issues tend to bubble up more frequently when  under load/stress. Generating stress for your tests can help to expose
these issues.

the **hack** directory contains 2 shell scripts that support stress testing your package libs. Usage:

```sh
# stress test without checking for race conditions
hack/stress.sh

# stress test while checking for race conditions
hack/stress-race.sh
```

### Security

Use a tool like [Go Sec](https://github.com/securego/gosec) to inspect code for security problems.

```sh
go get github.com/securego/gosec/cmd/gosec

#run gosec enabling tests and vendor files. They are ignored by default.
gosec -tests -vendor -fmt=json -out=results.json ./...
```

- Check out [Snyk's Vulnerability DB](https://snyk.io/vuln) for issues you should watchout for.
- Read through [OWASP's Secure Coding Practices](https://github.com/OWASP/Go-SCP) guide for Golang
- For more security inspiration look at [Awesome Golang Security](https://github.com/guardrailsio/awesome-golang-security)

Want a shortcut? Checkout hack/security.sh

```sh
hack/security.sh
```

### Error Handling

Error Handling will be handled in one of the 3 standard ways:

```go
// String based
err := errors.New("something bad happened")
```

```go
// format based
err :=  fmt.Errorf("something bad happened")
```

For a small set of errors you can use type errors. You can find a demo on the [Golang tour](https://tour.golang.org/methods/19)

```go
// Custom
type CustomError struct {
    Code int
    KeepGoing  bool
    Message string
}

func (ce CustomError) Error(){
    return fmt.Sprintf("%d:%t:%s", ce.Code, ce.KeepGoing, ce.Message)
}
//...
if err := Foo(); err != nil {
    switch e := err.(type) {
    case *CustomError:
        // Do something interesting with e.Line and e.Col.
    case *SomeOtherError:
        // Abort and file an issue.
    default:
        log.Println(e)
    }
}
```

If you dont't have a small set of errors or you dont know how many errors you will have, consider using behavior checks instead.
This pattern is a bit more future-proof.

Note: this is inspired by https://medium.com/@srfrog/i-wouldnt-recommend-using-type-checking-for-errors-b32accc77dd8

```go

type customError interface{
   BehaviorA() bool
}
type CustomError struct{}
func (e CustomError) BehaviorA() bool {
   return true
}
func (e CustomError) Error() string {
   return "something bad happened!"
}
// this func could apply to multiple types
func IsBehaviorA(e error) bool {
   f, ok := e.(customError)
   return ok && f.BehaviorA()
}
// specific to this error
func IsBehaviorB(e error) bool {
   _, ok := e.(CustomError)
   return ok
}
main() {
   err := caller()
   if IsBehaviorB(err) || IsBehaviorA(err) {
      // place imgination here
   }
}
```
