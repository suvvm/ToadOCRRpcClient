#!/bin/bash

protoc -I rpc/toad_ocr_engine/idl rpc/toad_ocr_engine/idl/toad_ocr_engine.proto --go_out=plugins=grpc:rpc/toad_ocr_engine/idl
sed -i '' 's/ClientConnInterface/ClientConn/g' rpc/toad_ocr_engine/idl/toad_ocr_engine.pb.go
sed -i '' 's/SupportPackageIsVersion6/SupportPackageIsVersion4/g' rpc/toad_ocr_engine/idl/toad_ocr_engine.pb.go
