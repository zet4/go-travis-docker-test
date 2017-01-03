#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover -coverprofile=profile.out -covermode=atomic --trace --race --progress $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done