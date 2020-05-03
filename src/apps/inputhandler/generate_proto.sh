#!/bin/bash

protodir=../../../proto-buffers

mkdir genproto

protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/demo.proto