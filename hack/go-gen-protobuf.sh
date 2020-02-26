#!/bin/bash

set -o errexit

protoc="${PROTOC:-protoc}"
include_dir="$1"
input="$2"
go_out="$3"
output="$4"

if [[ -z "$include_dir" || -z "$input" || -z "$go_out" || -z "$output" ]]; then
  echo >&2 "Usage: $0 <include_dir> <input.proto> <go_out_dir> <output.pb.go>"
  exit 1
fi

"$protoc" -I "$include_dir" "--go_out=$go_out" "$input" || exit $?

structs=()

mapfile -t structs < <(
  \grep -E -h -o '^type ([A-Z]\w+) struct' "$output" |
  \awk '{ print $2 }' |
  \sort
)

json_file="$go_out/json.go"
json_marshal_tpl="$(cat <<EOF

func (m *%s) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}
EOF
)"

for struct in "${structs[@]}"; do
  if ! \grep -q -E "^func \\(m \\*${struct}\\) MarshalJSON\\(\\)" "$json_file"; then
    printf -v json_marshal_src "$json_marshal_tpl" "$struct"
    echo "$json_marshal_src" >> "$json_file"
  fi
done
