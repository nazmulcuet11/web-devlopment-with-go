package controllers

import (
	"html/template"
	"net/http"
	"note-api/models"
	"note-api/repositories"

	"github.com/gorilla/mux"
)

type NoteTemplatleController struct {
	templates map[string]*template.Template
	repo      repositories.NoteRepository
}

func NewNoteTemplatleController(repo repositories.NoteRepository) *NoteTemplatleController {
	c := NoteTemplatleController{
		make(map[string]*template.Template),
		repo,
	}
	c.initializeTemplates()
	return &c
}

func (c *NoteTemplatleController) initializeTemplates() {
	c.templates["index"] = template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
	c.templates["add"] = template.Must(template.ParseFiles("templates/base.html", "templates/add.html"))
	c.templates["edit"] = template.Must(template.ParseFiles("templates/base.html", "templates/edit.html"))
}

func (c *NoteTemplatleController) renderTemplate(
	w http.ResponseWriter, name string,
	template string,
	viewModel interface{},
) {
	tmpl, ok := c.templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler for "/" which render the index page
func (c *NoteTemplatleController) GetNotes(w http.ResponseWriter, r *http.Request) {
	c.renderTemplate(w, "index", "base", c.repo.GetAllNotes())
}

// Handler for "/notes/add" for add a new item
func (c *NoteTemplatleController) AddNote(w http.ResponseWriter, r *http.Request) {
	c.renderTemplate(w, "add", "base", nil)
}

// Handler for "/notes/save" for save a new item into the data store
func (c *NoteTemplatleController) SaveNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	c.repo.AddNote(models.Note{Title: title, Description: description})
	http.Redirect(w, r, "/", http.StatusFound)
}

// Handler for "/notes/edit/{id}" to edit an existing item
func (c *NoteTemplatleController) EditNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	note, err := c.repo.GetNoteById(id)
	if err != nil {
		http.Error(w, "Could not find note", http.StatusBadRequest)
		return
	}
	c.renderTemplate(w, "edit", "base", note)
}

// Handler for "/notes/update/{id}" which update an item into the data store
func (c *NoteTemplatleController) UpdateNote(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	id := vars["id"]
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	noteToUpdate := models.Note{Id: id, Title: title, Description: description}
	err := c.repo.UpdateNote(noteToUpdate)
	if err != nil {
		http.Error(w, "Could not update note", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Handler for "/notes/delete/{id}" which delete an item form the store
func (c *NoteTemplatleController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := c.repo.DeleteNoteById(id)
	if err != nil {
		http.Error(w, "Could not delete note", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
