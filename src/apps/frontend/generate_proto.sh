#!/bin/bash

protodir=../../../proto-buffers

mkdir proto

protoc --go_out=plugins=grpc:proto -I $protodir $protodir/demo.proto