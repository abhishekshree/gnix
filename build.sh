#!/bin/sh

for i in "cat" "cp" "dirname" "echo" "grep" "ls" "mkdir" "pwd" "sleep" "touch" "whoami";
do
    printf "Compiling '$i'..."
    go build -o build/$i ./$i/${i}.go
    printf " Done\n"
done
