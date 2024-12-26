#!/bin/bash

cd $(dirname "$0")/..
go test github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity || FAILED=1

if [ -n "$FAILED" ]; then
    exit 1
fi
