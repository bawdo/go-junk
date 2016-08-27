package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)
	fmt.Println("Everything is set up!")
}
