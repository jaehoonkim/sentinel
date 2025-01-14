#!/usr/bin/env bash

pwd=$PWD



# Show env vars
grep -v '^#' .sentinel.env

# Export env vars
export $(grep -v '^#' .sentinel.env | xargs)

# path="$(dirname "$0")"
echo $pwd
# echo $path

cd ../sentinel

../cmd/repl.sh


cd $pwd