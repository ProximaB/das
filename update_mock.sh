# mockgen.sh batch updates all mock objects.
# To run this script:
# $ cd $GOPATH/src/github.com/DancesportSoftware/das
# $ sh update_mock.sh
#
# Copyright 2018 Yubing Hou. All rights reserved.
# Use of this source code is governed by GPL license
# that can be found in the LICENSE file
#! /bin/bash

# directory to mock to
MOCK_ROOT_DIR='./mock/'

# list of modules to mock
modules=('./businesslogic/' './businesslogic/reference/')

# remove existing mock and create a new one
rm -rf $MOCK_ROOT_DIR
mkdir $MOCK_ROOT_DIR

for m in "${modules[@]}"
do
    mkdir -p $MOCK_ROOT_DIR$m # create sub modules
    for each in $m*.go
    do
        if [[ "$each" != *"test.go"* ]]; then
            dest=${each:2:${#each}-5}
            mockgen -source=$each > $MOCK_ROOT_DIR/${dest}_mock.go
        fi
    done
done