#!/bin/sh

for i in $(ls *.go)
do
    printf "Compiling '$i'..."
    go build -o build/${i%.*} $i
    printf " Done\n"
done