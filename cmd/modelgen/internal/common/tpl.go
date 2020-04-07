package common

import (
	"fmt"
	"io"
	"text/template"
)

func CreateTemplate(name, source string, funcMaps ...template.FuncMap) *template.Template {
	t := template.New(name).Option("missingkey=error")

	for _, fmap := range funcMaps {
		t = t.Funcs(fmap)
	}

	return template.Must(t.Parse(source))
}

func ExecTemplate(w io.Writer, tpl *template.Template, vals map[string]interface{}) error {
	if err := tpl.Execute(w, vals); err != nil {
		return fmt.Errorf("error executing template %s: %w", tpl.Name(), err)
	}

	return nil
}
