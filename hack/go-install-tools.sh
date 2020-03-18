#!/bin/bash

set -o errexit

if [[ -z "$NS_ROOT" ]]; then
  echo "NS_ROOT env variable must be set" >&2
  exit 1
fi

if [[ -z "$NS_GO_TOOLS_PATH" ]]; then
  echo "NS_GO_TOOLS_PATH env variable must be set" >&2
  exit 1
fi

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

readarray -t tools < <(awk "$awkprog" "$NS_ROOT/tools.go")

export GOBIN="$NS_GO_TOOLS_PATH"

echo "Installing tools in $NS_GO_TOOLS_PATH"
echo "Make sure to add the directory to \$PATH"
echo

mkdir -p "$NS_GO_TOOLS_PATH"

for pkg in "${tools[@]}"; do
  echo " * $pkg"
  go install "$pkg"
done
