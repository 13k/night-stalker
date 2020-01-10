#!/bin/sh

pkgs="$(
  \grep '^'$'\t' go.mod |
  \grep -v '// indirect' |
  \awk '{ print $1 }'
)"

for pkg in $pkgs; do
  printf -- "-----\n%s\n-----\n" "$pkg"
  go get "$pkg"
done
