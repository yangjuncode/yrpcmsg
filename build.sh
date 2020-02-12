
#!/bin/bash

mkdir -p dist
pbjs --no-create --no-verify --no-convert --no-delimited --force-long   -t static-module -w es6 -o dist/index.js yrpcmsg.proto
pbts -o dist/index.d.ts dist/index.js

protoc --gogofaster_out=. yrpcmsg.proto
rm yrpcmsg.pb.go
mv github.com/yangjuncode/yrpcmsg/yrpcmsg.pb.go .
rm -rf github.com
