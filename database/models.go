// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Note struct {
	ID      int64
	Title   string
	Content sql.NullString
}