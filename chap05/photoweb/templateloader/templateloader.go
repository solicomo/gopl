package templateloader

import (
	"errors"
	"html/template"
	"io/ioutil"
	"path"
)

func LoadTemplates(dir, ext string) (tpl *template.Template, err error) {
	tpl = template.New("")

	if len(dir) > 0 && dir[len(dir)-1] != '/' {
		dir = dir + "/"
	}

	err = loadTemplates(dir, ext, tpl)
	return
}

func loadTemplates(dir, ext string, tpl *template.Template) (err error) {
	if tpl == nil {
		err = errors.New("create new template failed.")
		return
	}

	pName := tpl.Name()
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
			cTpl  := tpl.New(cName)
			er := loadTemplates(fName + "/", ext, cTpl)
			if er != nil {
				err = er
				return
			}

			continue
		}

		if ex := path.Ext(fName); ex != ext {
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
