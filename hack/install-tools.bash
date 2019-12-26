#!/bin/bash

set -o errexit

ROOT="$(realpath "$(dirname "${BASH_SOURCE[0]}")/..")"

declare -a tools

awkprog="$(
  cat <<-'EOF'
/^\s+_ ".+"$/ {
  s = $2
  gsub("\"", "", s)
  print s
}
EOF
)"

readarray -t tools < <(awk "$awkprog" "$ROOT/tools.go")

export GOBIN="$ROOT/bin"

echo "Installing tools in $GOBIN"
echo "Make sure to add the directory to \$PATH"
echo

mkdir -p "$GOBIN"

for pkg in "${tools[@]}"; do
  echo " * $pkg"
  go install "$pkg"
done
