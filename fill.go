package nut

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func FillTemplate(tpl *template.Template, data Structure, saveFile string) error {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("failed to fill template:%v", err)
	}

	fd, err := os.OpenFile(saveFile, os.O_APPEND, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file %s:%v", saveFile, err)
	}

	_, err = fd.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write file %s:%v", saveFile, err)
	}

	return nil
}
