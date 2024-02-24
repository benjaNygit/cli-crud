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

	// Type Note
	r.HandleFunc("/types", routers.GetTypesNoteHandler).Methods("GET")
	r.HandleFunc("/types", routers.PostTypeNoteHandler).Methods("POST")
	r.HandleFunc("/types/{id}", routers.GetTypeNoteHandler).Methods("GET")
	r.HandleFunc("/types/{id}", routers.DeleteTypeNoteHandler).Methods("DELETE")

	// Note
	r.HandleFunc("/notes", routers.GetNotesHandler).Methods("GET")
	r.HandleFunc("/notes", routers.PostNotesHandler).Methods("POST")
	r.HandleFunc("/notes/{id}", routers.GetNoteHandler).Methods("GET")
	r.HandleFunc("/notes/{id}", routers.DeleteNoteHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
