package handler

import (
	"fmt"
	"net/http"
)

type Order struct {}

func (o *Order) UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create order")
}

func (o *Order) GetFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list orders")
}