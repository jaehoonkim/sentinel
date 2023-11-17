#!/usr/bin/env bash

pwd=$PWD



# Show env vars
grep -v '^#' .synapse.env

# Export env vars
export $(grep -v '^#' .synapse.env | xargs)

# path="$(dirname "$0")"
echo $pwd
# echo $path

cd ../synapse

../cmd/repl.sh


cd $pwd