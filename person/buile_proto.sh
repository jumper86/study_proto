#!/bin/bash

protoc --proto_path=. --go_out=. --go-grpc_out=. code/*.proto