#!/usr/bin/env bash

set -e

cd $(dirname $0)

for f in *.proto; do
  echo -n "Checking $f..."
  protoc $f -o tmp.protoset -I . -I ../../internal/testdata
  echo "  good"
done

rm tmp.protoset
