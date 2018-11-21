#!/usr/bin/env bash
# mockgen.sh batch updates all mock objects.
# To run this script:
# $ cd $GOPATH/src/github.com/DancesportSoftware/das
# $ bash update_mock.sh
# On Windows, run this script in a Linux subsystem, or embedded Bash
#
# Dancesport Application System (DAS)
# Copyright (C) 2018 Yubing Hou
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
# directory to mock to
MOCK_ROOT_DIR="./mock/"
#
# list of modules to mock
declare -a modules=("./businesslogic/" "./businesslogic/reference/")
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
