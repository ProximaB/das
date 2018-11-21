#!/usr/bin/env bash

# test script for codecov
#
# Troubleshoot for Windows:
# This program may not be able to run on Windows bash since Windows uses a different file-encoding from ASCII
# To remove non-ASCII characters from this script, run:
# $ hexdump -C go.test.sh
# Look for 0d 0a sqeuences. We can strip thie \r (0d) with the tr command:
# $ cat go.test.sh | tr -d '\r' >> go.test2.sh
# Then copy the content from go.test2.sh to go.test.sh, and remove go.test2.sh
# This solution was from Stackoverflow:
# https://stackoverflow.com/questions/5491634/shell-script-error-expecting-do

# setup environment variable
export POSTGRES_CONNECTION="host\=127.0.0.1 user\=dasdev password\=dAs\!@#\$1234 dbname\=das sslmode\=disable"
export FIREBASE_SERVICE_ACCOUNT_KEY="$GOPATH/src/github.com/DancesportSoftware/das/secret/firebase-auth-dev-02.json"

set -e
echo "" > coverage.txt
#
for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
