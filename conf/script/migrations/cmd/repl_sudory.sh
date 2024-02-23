#!/usr/bin/env bash

pwd=$PWD



# Show env vars
grep -v '^#' .morpheus.env

# Export env vars
export $(grep -v '^#' .morpheus.env | xargs)

# path="$(dirname "$0")"
echo $pwd
# echo $path

cd ../morpheus

../cmd/repl.sh


cd $pwd