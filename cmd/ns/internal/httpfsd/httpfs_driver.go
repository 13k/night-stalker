package httpfsd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/golang-migrate/migrate/v4/source"
)

type httpfsDriver struct {
	migrations *source.Migrations
	fs         http.FileSystem
	path       string
}

func New(fs http.FileSystem, path string) (source.Driver, error) {
	root, err := fs.Open(path)

	if err != nil {
		return nil, err
	}

	files, err := root.Readdir(0)

	if err != nil {
		_ = root.Close()
		return nil, err
	}

	if err = root.Close(); err != nil {
		return nil, err
	}

	ms := source.NewMigrations()

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		m, err := source.DefaultParse(file.Name())

		if err != nil {
			continue // ignore files that we can't parse
		}

		if !ms.Append(m) {
			return nil, fmt.Errorf("duplicate migration file: %v", file.Name())
		}
	}

	driver := &httpfsDriver{
		fs:         fs,
		path:       path,
		migrations: ms,
	}

	return driver, nil
}

func (h *httpfsDriver) Open(url string) (source.Driver, error) {
	return nil, fmt.Errorf("not implemented")
}

func (h *httpfsDriver) Close() error {
	return nil
}

func (h *httpfsDriver) First() (uint, error) {
	if v, ok := h.migrations.First(); ok {
		return v, nil
	}

	return 0, &os.PathError{
		Op:   "first",
		Path: h.path,
		Err:  os.ErrNotExist,
	}
}

func (h *httpfsDriver) Prev(version uint) (uint, error) {
	if v, ok := h.migrations.Prev(version); ok {
		return v, nil
	}

	return 0, &os.PathError{
		Op:   fmt.Sprintf("prev for version %v", version),
		Path: h.path,
		Err:  os.ErrNotExist,
	}
}

func (h *httpfsDriver) Next(version uint) (uint, error) {
	if v, ok := h.migrations.Next(version); ok {
		return v, nil
	}

	return 0, &os.PathError{
		Op:   fmt.Sprintf("next for version %v", version),
		Path: h.path,
		Err:  os.ErrNotExist,
	}
}

func (h *httpfsDriver) ReadUp(version uint) (io.ReadCloser, string, error) {
	m, ok := h.migrations.Up(version)

	if !ok {
		return nil, "", &os.PathError{
			Op:   fmt.Sprintf("read version %v", version),
			Path: h.path,
			Err:  os.ErrNotExist,
		}
	}

	body, err := h.fs.Open(path.Join(h.path, m.Raw))

	if err != nil {
		return nil, "", err
	}

	return body, m.Identifier, nil
}

func (h *httpfsDriver) ReadDown(version uint) (io.ReadCloser, string, error) {
	m, ok := h.migrations.Down(version)

	if !ok {
		return nil, "", &os.PathError{
			Op:   fmt.Sprintf("read version %v", version),
			Path: h.path,
			Err:  os.ErrNotExist,
		}
	}

	body, err := h.fs.Open(path.Join(h.path, m.Raw))

	if err != nil {
		return nil, "", err
	}

	return body, m.Identifier, nil
}
