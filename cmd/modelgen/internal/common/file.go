package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenFilename(srcFile, suffix string) string {
	ext := filepath.Ext(srcFile)
	base := strings.TrimSuffix(srcFile, ext)
	return fmt.Sprintf("%s%s%s", base, suffix, ext)
}

func CreateFile(fname string, code []byte) error {
	file, err := os.Create(fname)

	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", fname, err)
	}

	defer file.Close()

	if _, err = file.Write(code); err != nil {
		return fmt.Errorf("error writing to file %s: %w", fname, err)
	}

	return nil
}
