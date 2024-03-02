package template

import (
	"html/template"
	"net/http"
	"note-api/models"
	"note-api/repositories"

	"github.com/gorilla/mux"
)

var templates map[string]*template.Template

func InitializeTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/base.html", "templates/add.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/base.html", "templates/edit.html"))
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler for "/" which render the index page
func GetNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", repositories.GetAllNotes())
}

// Handler for "/notes/add" for add a new item
func AddNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

// Handler for "/notes/save" for save a new item into the data store
func SaveNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	repositories.AddNote(models.Note{Title: title, Description: description})
	http.Redirect(w, r, "/", http.StatusFound)
}

// Handler for "/notes/edit/{id}" to edit an existing item
func EditNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	note, err := repositories.GetNoteById(id)
	if err != nil {
		http.Error(w, "Could not find note", http.StatusBadRequest)
		return
	}
	renderTemplate(w, "edit", "base", note)
}

// Handler for "/notes/update/{id}" which update an item into the data store
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	id := vars["id"]
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	noteToUpdate := models.Note{Id: id, Title: title, Description: description}
	err := repositories.UpdateNote(noteToUpdate)
	if err != nil {
		http.Error(w, "Could not update note", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Handler for "/notes/delete/{id}" which delete an item form the store
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := repositories.DeleteNoteById(id)
	if err != nil {
		http.Error(w, "Could not delete note", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
