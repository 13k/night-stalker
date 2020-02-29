#!/bin/bash

set -o errexit

protoc="${PROTOC:-protoc}"
include_dir="$1"
input="$2"
go_out="$3"

if [[ -z "$include_dir" || -z "$input" || -z "$go_out" ]]; then
  echo >&2 "Usage: $0 <include_dir> <input.proto> <go_out_dir>"
  exit 1
fi

exec "$protoc" \
  -I "$include_dir" \
  "--go_out=paths=source_relative:$go_out" \
  "--go-json_out=enums_as_ints,emit_defaults,orig_name:$go_out" \
  "$input"
