package routers

import (
	"encoding/json"
	"net/http"

	"github.com/benjaNygit/cli-crud/go/notes/db"
	"github.com/benjaNygit/cli-crud/go/notes/models"
	"github.com/gorilla/mux"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	db.DB.Find(&notes)

	json.NewEncoder(w).Encode(&notes)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var note models.Note

	db.DB.First(&note, params["id"])
	if note.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Note Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&note)
}

func PostNotesHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)

	object := db.DB.Create(&note)
	err := object.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&note)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

}
