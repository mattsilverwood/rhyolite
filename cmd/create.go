package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattsilverwood/rhyolite/database"
)

func CreateNote(ctx context.Context, dbconn *sql.DB, newNote database.CreateNoteParams) (database.Note, error) {
	query := database.New(dbconn)
	createdNote, err := query.CreateNote(ctx, newNote)

	if err != nil {
		return database.Note{}, errors.New(fmt.Sprintf("Could not create note: %s", err))
	}

	return createdNote, nil

}
