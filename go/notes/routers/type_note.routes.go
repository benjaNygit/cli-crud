package routers

import (
	"encoding/json"
	"net/http"

	"github.com/benjaNygit/cli-crud/go/notes/db"
	"github.com/benjaNygit/cli-crud/go/notes/models"
	"github.com/gorilla/mux"
)

func GetTypesNoteHandler(w http.ResponseWriter, r *http.Request) {
	var types_note []models.TypeNote
	db.DB.Find(&types_note)

	json.NewEncoder(w).Encode(&types_note)
}

func GetTypeNoteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var type_note models.TypeNote

	db.DB.First(&type_note, params["id"])
	if type_note.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Type Note Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&type_note)
}

func PostTypeNoteHandler(w http.ResponseWriter, r *http.Request) {
	var type_note models.TypeNote
	json.NewDecoder(r.Body).Decode(&type_note)
	object := db.DB.Create(&type_note)
	err := object.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Type Note Bad Request"))
		return
	}

	json.NewEncoder(w).Encode(&type_note)
}

func DeleteTypeNoteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var type_note models.TypeNote

	db.DB.First(&type_note, params["id"])
	if type_note.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Type Note Delete Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&type_note)
	w.WriteHeader(http.StatusOK)
}
