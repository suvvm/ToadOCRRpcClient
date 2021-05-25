#!/bin/bash

protoc -I rpc/toad_ocr_preprocessor/idl rpc/toad_ocr_preprocessor/idl/toad_ocr_preprocessor.proto --go_out=plugins=grpc:rpc/toad_ocr_preprocessor/idl
sed -i '' 's/ClientConnInterface/ClientConn/g' rpc/toad_ocr_preprocessor/idl/toad_ocr_preprocessor.pb.go
sed -i '' 's/SupportPackageIsVersion6/SupportPackageIsVersion4/g' rpc/toad_ocr_preprocessor/idl/toad_ocr_preprocessor.pb.go
