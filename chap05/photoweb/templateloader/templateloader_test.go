package templateloader

import (
	"fmt"
	"html/template"
	"testing"
)

const (
	TPL_DIR = "./tpls"
	TPL_EXT = ".tpl"
)

func dumpTemplates(tpls *template.Template) {
	for _, t := range tpls.Templates() {
		fmt.Println(t.Name())
	}
}

func TestLoadTemplates(t *testing.T) {
	tpls, err := LoadTemplates(TPL_DIR, TPL_EXT)

	if err != nil {
		t.Error(err)
	}

	dumpTemplates(tpls)
}
