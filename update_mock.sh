# mockgen.sh batch updates all mock objects
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
#    echo Creating mock objects for module $SOURCE_DIR$m
#    if [ -f $SOURCE_DIR$m ]; then # a regular go file
#        echo \tCreating mock object for $SOURCE_DIR$m
#        mkdir -p $MOCK_ROOT_DIR/$(dirname "$m")
#        if [[$SOURCE_DIR$m -ne *"_test.go"]]; then
#
#        fi
#    elif [ -d $SOURCE_DIR$m ]; then # a regular directory

done

# for filename in ./businesslogic/*.go; do
#    mockgen -source="$filename" > ./mock/businesslogic/$(basename "$filename" .go)_mock.go
# done

# for filename in ./businesslogic/reference/*.go; do
#    mockgen -source="$filename" > ./mock/businesslogic/reference/$(basename "$filename" .go)_mock.go
# done