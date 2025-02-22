package handlers

import (
	"SimpleGoProject/utils"
	"fmt"
	"net/http"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new book")
}

func ShowBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Show the details of book %d\n", id)
}
