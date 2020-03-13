package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/upload", upload)
	http.HandleFunc("/download", download)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}

}

func upload(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	file, fileHead, err := req.FormFile("file")
	if err != nil {
		io.WriteString(w, "Read File Error")
		return
	}

	defer file.Close()
	log.Println("filename: " + fileHead.Filename)

	newFile, err := os.Create(fileHead.Filename)
	if err != nil {
		io.WriteString(w, "Create File Error")
		return
	}

	defer newFile.Close()

	if _, err = io.Copy(newFile, file); err != nil {
		io.WriteString(w, "Write File Error")
		return
	}

	io.WriteString(w, "Upload File Succeed")

}

func download(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	filename := req.FormValue("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request")
		return
	}

	log.Println("download file name => ", filename)

	file, err := os.Open(filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "file not found")
		return
	}

	defer file.Close()

	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment;filename=\""+filename+"\"")
	if _, err = io.Copy(w, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Write File Error")
		return
	}
}
