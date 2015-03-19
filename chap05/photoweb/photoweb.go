package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads/"
)

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
	w.Header().Set("Content-Type", "image")

	http.ServeFile(w, r, image)
}


func main() {
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/view", handleView)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
