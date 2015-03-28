package templateloader

import (
	"errors"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

func LoadTemplates(dir, pattern string) (tpl *template.Template, err error) {
	tpl = template.New("")

	if len(dir) > 0 && dir[len(dir)-1] != '/' {
		dir = dir + "/"
	}

	err = loadTemplates(dir, pattern, tpl.Name(), tpl)
	return
}

func loadTemplates(dir, pattern, pName string, tpl *template.Template) (err error) {
	if tpl == nil {
		err = errors.New("create new template failed.")
		return
	}

	if len(pName) > 0 && dir[len(pName)-1] != '/' {
		pName += "/"
	}

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, fi := range fis {
		cName := pName + fi.Name()
		fName := dir + fi.Name()

		if fi.IsDir() {
			er := loadTemplates(fName+"/", pattern, cName, tpl)
			if er != nil {
				err = er
				return
			}

			continue
		}

		matched, er := filepath.Match(pattern, fi.Name())
		if er != nil {
			err = er
			return
		}

		if !matched {
			continue
		}

		b, er := ioutil.ReadFile(fName)
		if er != nil {
			err = er
			return
		}

		s := string(b)

		cTpl := tpl.New(cName)

		cTpl, er = cTpl.Parse(s)
		if er != nil {
			err = er
			return
		}
	}

	return
}
