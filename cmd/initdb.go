package cmd

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

func InitDB(ctx context.Context, dbconn *sql.DB, schema string) error {

	// Initialise database by running schema
	if _, err := dbconn.ExecContext(ctx, schema); err != nil {
		return errors.New(fmt.Sprintf("Could not apply schema: %s", err))
	}

	return nil
}
