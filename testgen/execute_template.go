package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"text/template"
)

func ExecTemplate(tplName, tplFile, output string, data any) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New(tplName).Parse(string(tpl))
	if err != nil {
		return fmt.Errorf("error parsing template %s: %s", tplFile, err)
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, data); err != nil {
		return err
	}

	formattedCode, err := format.Source(code.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(output, formattedCode, 0644); err != nil {
		return err
	}

	return nil
}
