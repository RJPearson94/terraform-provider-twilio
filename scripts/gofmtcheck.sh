#!/usr/bin/env bash

gofmt_files=$(find . -name '*.go' | grep -v vendor | xargs gofmt -l -s)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

exit 0