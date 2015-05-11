package main

import (
	"fmt"
	"./templateloader"
	//"gopl/chap05/photoweb/templateloader"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"html/template"
)

const (
	PUBLIC_DIR   = "./public/"
	UPLOAD_DIR   = "./public/upload/"
	STATIC_DIR   = "./public/static/"
	TEMPLATE_DIR = "./tpl/"
)

var gTpls *template.Template = nil

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload.html", nil)
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
		check(err)

		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)

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

	renderHtml(w, "view.html", id)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fis, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)

	images := []string{}
	for _, fi := range fis {
		if fi.IsDir() || fi.Name()[0] == '.' {
			continue
		}

		images = append(images, fi.Name())
	}

	renderHtml(w, "index.html", images)
}

func renderHtml(w http.ResponseWriter, tpl string, data interface{}) {
	check(gTpls.ExecuteTemplate(w, tpl, data))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				renderHtml(w, "50x.html", err)
			}
		}()

		fn(w, r)
	}
}

func main() {
	var err error
	gTpls, err = templateloader.LoadTemplates(TEMPLATE_DIR, "*.html")
	if err != nil {
		fmt.Println("load templates failed:", err)
		return
	}

	http.HandleFunc("/", safeHandler(handleIndex))
	http.HandleFunc("/upload", safeHandler(handleUpload))
	http.HandleFunc("/view", safeHandler(handleView))
	http.Handle("/upload/", http.FileServer(http.Dir(PUBLIC_DIR)))
	http.Handle("/static/", http.FileServer(http.Dir(PUBLIC_DIR)))
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
