#! /bin/sh

find . -name *.proto |
while read filename
do
    echo "proto-compiler $filename running...."
    protoc --micro_out=. --go_out=. $filename
    echo "proto-compiler $filename done"
done