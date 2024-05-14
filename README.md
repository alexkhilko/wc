## About

Implementation of Unix `wc` tool written in Go

## Build

Do `make build` to create the executable. It will be saved under `bin` directory.

## Run

Get the number of bytes
`./bin/wc -c <path_to_file>`

Get the number of line breaks
`./bin/wc -l <path_to_file>`

Get the number of characters
`./bin/wc -m <path_to_file>`

Get the number of words
`./bin/wc -w <path_to_file>`

