#!/usr/bin/env bash -e

# This script is inspired by Dave Cheney's Blog
# post https://dave.cheney.net/2013/06/19/stress-test-your-go-packages

# Quote from the blog:
# > Concurrency or memory correctness errors are more likely to show up at
# > higher concurrency levels (higher values of GOMAXPROCS). I use this script
# > when testing my packages, or when reviewing code that goes into the
# > standard library.

# ID the working directory.
#################
WORKING=generated

# Find all of the tests that are NOT integration tests, isolating their directories
#################
# Unix flavored
# PKGS=(`find . -type f \( -name "*test.go" ! -iname "*integration*" \)  | sed -r 's|/[^/]+$||' |sort |uniq`)

# Mac flavored
PKGS=(`find . -type f \( -name "*test.go" ! -iname "*integration*" \)  | sed -E 's|/[^/]+$||' |sort -u`)

# Compile the tests for each package
# go test -c
for PKG in "${PKGS[@]}"
do
    TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')
    PKG_BASE=$(basename $PKG)
    echo ""
    echo "-----------------------------"
    echo $PKG_BASE
    echo "-----------------------------"

    # echo "go test -c $PKG"
    go test -c $PKG

    # echo "go test -benchmem -memprofile=$TIMESTAMP.$PKG_BASE.mem.out -cpuprofile=$TIMESTAMP.$PKG_BASE.cpu.out -trace=$TIMESTAMP.$PKG_BASE.trace.out -bench=.  $PKG > $TIMESTAMP.$PKG_BASE.benchmark.txt"
    go test -benchmem -memprofile=$TIMESTAMP.$PKG_BASE.mem.out -cpuprofile=$TIMESTAMP.$PKG_BASE.cpu.out -trace=$TIMESTAMP.$PKG_BASE.trace.out -bench=.  $PKG > $TIMESTAMP.$PKG_BASE.benchmark.txt

    mkdir -p $WORKING/$PKG_BASE
    mv *$PKG_BASE* $WORKING/$PKG_BASE
    echo ""
    echo "Don't forget to investigate $PKG_BASE test results using pprof"
    echo "go tool pprof -http=localhost:9999 $WORKING/$PKG_BASE/$PKG_BASE.test $WORKING/$PKG_BASE/$TIMESTAMP.$PKG_BASE.cpu.out"
    echo "go tool pprof -http=localhost:9999 -alloc_objects $WORKING/$PKG_BASE/$PKG_BASE.test $WORKING/$PKG_BASE/$TIMESTAMP.$PKG_BASE.mem.out"
    echo "go tool trace -http=localhost:9999 $WORKING/$PKG_BASE/$PKG_BASE.test $WORKING/$PKG_BASE/$TIMESTAMP.$PKG_BASE.trace.out"
    echo "======== END $PKG_BASE ========"

done

echo ""
echo "-----------------------------"
echo "Addition Tips"
echo "-----------------------------"
echo "Additonal flags to use when investigating memory profile data (You can only use one of them at a time)"
echo "  -inuse_space      Display in-use memory size"
echo "  -inuse_objects    Display in-use object counts"
echo "  -alloc_space      Display allocated memory size"
echo "  -alloc_objects    Display allocated object counts"
