package test_data

import (
	"note-api/models"
	"note-api/repositories"
	"time"
)

var TestNotes = []models.Note{
	{Id: "1", Title: "Note 1", Description: "Note 1 description", CreatedOn: time.Now()},
	{Id: "2", Title: "Note 2", Description: "Note 2 description", CreatedOn: time.Now()},
	{Id: "3", Title: "Note 3", Description: "Note 3 description", CreatedOn: time.Now()},
	{Id: "4", Title: "Note 4", Description: "Note 4 description", CreatedOn: time.Now()},
	{Id: "5", Title: "Note 5", Description: "Note 5 description", CreatedOn: time.Now()},
}

func PopulateTestNotes() {
	repositories.AddNotes(TestNotes)
}
