#!/bin/sh

OUTPUT_DIR=bin

mkdir -p $OUTPUT_DIR


go build -o $OUTPUT_DIR ./cmd/...