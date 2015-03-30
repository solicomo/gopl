package main

import (
	"fmt"
	"gopl/chap05/photoweb/templateloader"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"html/template"
)

const (
	UPLOAD_DIR   = "./upload/"
	STATIC_DIR   = "./static/"
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

	renderHtml(w, "view.html", id)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fis, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	images := []string{}
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		images = append(images, fi.Name())
	}

	renderHtml(w, "upload.html", images)
}

func renderHtml(w http.ResponseWriter, tpl string, data interface{}) {
	err := gTpls.ExecuteTemplate(w, tpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var err error
	gTpls, err = templateloader.LoadTemplates(TEMPLATE_DIR, "*.html")
	if err != nil {
		fmt.Println("load templates failed:", err)
		return
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/view", handleView)
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir(UPLOAD_DIR))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR))))
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
