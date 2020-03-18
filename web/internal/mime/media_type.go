package mime

import (
	"mime"
)

type MediaType struct {
	Type   string
	Params map[string]string

	serialized string
	exts       []string
}

func ParseMediaType(v string) (*MediaType, error) {
	if v == "" {
		return nil, nil
	}

	t, params, err := mime.ParseMediaType(v)

	if err != nil {
		return nil, err
	}

	mt := &MediaType{
		Type:   t,
		Params: params,
	}

	if _, err := mt.Extensions(); err != nil {
		return nil, err
	}

	return mt, nil
}

func (t *MediaType) Serialize() string {
	if t.serialized == "" {
		t.serialized = mime.FormatMediaType(t.Type, t.Params)
	}

	return t.serialized
}

func (t *MediaType) Extensions() ([]string, error) {
	if t.exts == nil {
		var err error

		t.exts, err = mime.ExtensionsByType(t.Type)

		if err != nil {
			return nil, err
		}
	}

	return t.exts, nil
}
