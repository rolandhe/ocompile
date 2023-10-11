#!/bin/sh

mkdir go
antlr4 -Dlanguage=Go Service.g4 -o ./go/
cp ./go/*.go ../parser/
rm -fr go
