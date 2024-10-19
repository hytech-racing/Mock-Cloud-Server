package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type File struct {}

func (o *File) UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploaded file successfully")

	response := map[string]string{"message": "success"}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (o *File) GetFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list orders")
}