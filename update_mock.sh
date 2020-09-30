#!/usr/bin/env bash
# mockgen.sh batch updates all mock objects.
# To run this script:
# $ cd $GOPATH/src/github.com/ProximaB/das
# $ bash update_mock.sh
# On Windows, run this script in a Linux subsystem, or embedded Bash
#
# Dancesport Application System (DAS)
# Copyright (C) 2018 Yubing Hou
#
# directory to mock to
MOCK_ROOT_DIR="./mock/"
#
# list of modules to mock
declare -a modules=("./businesslogic/")
#
# remove existing mock and create a new one
rm -rf $MOCK_ROOT_DIR
mkdir $MOCK_ROOT_DIR
#
for m in "${modules[@]}";
do
    mkdir -p $MOCK_ROOT_DIR$m # create sub modules
    for each in $m*.go;
    do
        if [[ "$each" != *"test.go"* ]]; then
            dest=${each:2:${#each}-5}
            mockgen -source=$each > $MOCK_ROOT_DIR/${dest}_mock.go
        fi
    done
done
