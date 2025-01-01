package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattsilverwood/rhyolite/database"
)

func ListAllNotes(ctx context.Context, dbconn *sql.DB) ([]database.Note, error) {
	query := database.New(dbconn)
	allnotes, err := query.ListAllNotes(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not complete query: %s", err))
	}

	return allnotes, nil

}
