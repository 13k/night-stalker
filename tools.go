// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/markbates/pkger/cmd/pkger"
	_ "github.com/mitchellh/protoc-gen-go-json"
)
