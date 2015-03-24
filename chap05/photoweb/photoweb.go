package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./upload/"
	STATIC_DIR = "./static/"
	TEMPLATE_DIR = "./tpl/"
)

func loadTpls(dir, ext string, tpl *template.Template) (err error) {
	if tpl == nil {
		err = errors.New("invalid template")
		return
	}

	pName := tpl.Name()

	if len(pName) > 0 && pName[len(pName)-1] != '/' {
		pName = pName + "/"
	}

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, fi := range fis {
		cName := pName + fi.Name()
		ct := tpl.New(cName)

		if fi.IsDir() {
			er := loadTpls(dir + "/" + name, ext, ct)
			if er != nil {
				err = er
				return
			}

			//TODO:
			continue
		}

		if ex := path.Ext(name); ex != ext {
			continue
		}

		ct, er := ct.ParseFiles(cName)
		if er != nil {
			err = er
			return
		}
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, `
		<html>
		<head><title>Upload</title></head>
		<body>
		<form method="POST" action="/upload" enctype="multipart/form-data">
		Choose an image to upload: <input name="image" type="file" />
		<input type="submit" value="Upload" />
		</form>
		</body>
		</html>`)

		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)

		return
	}

	http.Error(w, r.Method+" is not allowed", http.StatusMethodNotAllowed)
}

func handleView(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if len(id) == 0 {
		http.Error(w, "please specify an ID.", http.StatusNotFound)
		return
	}

	image := UPLOAD_DIR + id
	if exists := isExists(image); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")

	http.ServeFile(w, r, image)
}

func isExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fis, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `<html><head><title>Index</title></head><body><a href="/upload">Upload</a><br /><ol>`
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		id := fi.Name()
		html += `<li><a href="/view?id=` + id + `">` + id + "</a></li>"
	}

	html += "</ol></body></html>"

	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/view", handleView)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

