package db

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func reflectSlicePtrLen(s interface{}) int {
	v := reflect.ValueOf(s)

	if k := v.Kind(); k != reflect.Ptr {
		return -1
	}

	v = reflect.Indirect(v)

	if k := v.Kind(); k != reflect.Slice {
		return -1
	}

	return v.Len()
}

func genSavepointName() string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	return fmt.Sprintf("nsdb_%s", id)
}
