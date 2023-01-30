#!/bin/bash
set -e

SCRIPTS_DIR=$(dirname "$0")
cd ${SCRIPTS_DIR}/../
PROJECT_DIR=$(pwd);
PROTO_DIR=${PROJECT_DIR}/proto

# Generate the Go code
buf generate proto

# Vault/Auth
# rm -rf internal/gen/core
# cp -r internal/gen/* .

# Blockchain
cp -r github.com/sonrhq/core/* .
rm -rf github.com
go mod tidy