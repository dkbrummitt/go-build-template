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
WORKING=$(pwd)
echo WORKING $WORKING

# Find all of the tests that are NOT integration tests, isolating their directories
#################
# Unix flavored
# PKGS=(`find . -type f \( -name "*test.go" ! -iname "*integration*" \)  | sed -r 's|/[^/]+$||' |sort |uniq`)

# Mac flavored
PKGS=(`find . -type f \( -name "*test.go" ! -iname "*integration*" \)  | sed -E 's|/[^/]+$||' |sort -u`)

echo PKGS $PKGS

# Compile the tests for each package
# go test -c
for PKG in "${PKGS[@]}"
do
    echo PKG $PKG
    go test -c -race $PKG

    # run the stress test for 60 seconds
    PKG_BASE=$(basename $PKG)
    echo PKG_BASE $PKG_BASE

    END=$((SECONDS+60))
    while [ $SECONDS -lt $END ] ; do
            export GOMAXPROCS=$[ 1 + $[ RANDOM % 128 ]]
            ./$PKG_BASE.test $@ 2>&1
    done

done
