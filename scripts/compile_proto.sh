#!/bin/sh

BASE_DIR="proto"
PROJECT_NAME="vl-account"
FILE_NAMES=""

for FILE_NAME in $(find "$BASE_DIR" -type f -name "*.proto"); do
  FILE_NAMES="${FILE_NAMES} ${FILE_NAME}"
done

rm -rf ./proto_generated/

if [ ! -z FILE_NAMES ]; then
  protoc --go_out=./ \
    --go-grpc_out=./ \
    --go_opt=module=$PROJECT_NAME \
    --go-grpc_opt=module=$PROJECT_NAME \
    $FILE_NAMES
fi
