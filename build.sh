#! /bin/sh

find . -name *.proto |
while read filename
do
    protoc --micro_out=. --go_out=. ./account/account.proto $filename    
done