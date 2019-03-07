#!/usr/bin/env bash

protoc --proto_path=pb --go_out=plugins=grpc:pb todo-service.proto