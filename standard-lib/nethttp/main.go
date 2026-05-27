package main

import (
	"encoding/json"
	"net/http"
)

type Notes struct{
	Title string `json:"title"`
	Description string `json:"description"`
}

var notes = make([]Notes, 0)

func getNotes(w http.ResponseWriter) []Notes {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
	return notes
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note Notes
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	notes = append(notes, note)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func deleteNoteById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	for i, note := range notes {
		if note.Title == id {
			notes = append(notes[:i], notes[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Note not found", http.StatusNotFound)
}

func notesHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		getNotes(w)
	case "POST":
		createNote(w, r)
	case "DELETE":
		deleteNoteById(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}




func main(){
	notes=append(notes,Notes{Title:"Note 1",Description:"This is note 1"})
	http.HandleFunc("/api/notes",notesHandler)
	http.ListenAndServe(":8080",nil)
	


}