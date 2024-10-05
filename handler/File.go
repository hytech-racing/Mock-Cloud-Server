package handler

import (
	"fmt"
	"net/http"
)

type File struct {}

func (o *File) UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploaded file successfully");
	w.WriteHeader(http.StatusOK);
}

func (o *File) GetFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list orders")
}