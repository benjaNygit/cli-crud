package main

import (
	"net/http"

	"github.com/benjaNygit/cli-crud/go/notes/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routers.HomeHandler)

	http.ListenAndServe(":3000", r)
}
