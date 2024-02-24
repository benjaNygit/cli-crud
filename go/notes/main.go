package main

import (
	"net/http"

	"github.com/benjaNygit/cli-crud/go/notes/db"
	"github.com/benjaNygit/cli-crud/go/notes/models"
	"github.com/benjaNygit/cli-crud/go/notes/routers"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.TypeNote{})
	db.DB.AutoMigrate(models.Note{})

	r := mux.NewRouter()
	r.HandleFunc("/", routers.HomeHandler)

	http.ListenAndServe(":3000", r)
}
