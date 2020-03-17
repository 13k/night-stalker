// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/markbates/pkger/cmd/pkger"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
