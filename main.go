package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/mattsilverwood/rhyolite/cmd"
	"github.com/mattsilverwood/rhyolite/database"
	_ "modernc.org/sqlite"
)

//go:embed database/schema.sql
var dbschema string

func main() {
	ctx := context.Background()
	dbconn, err := sql.Open("sqlite", "db")

	if err != nil {
		log.Fatalf("FATAL: Could not open database: %s", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("App not implemented.")
	}

	switch os.Args[1] {
	case "init":
		if len(os.Args) != 2 {
			log.Fatalf("Invalid usage. %s init", os.Args[0])
		}
		err := cmd.InitDB(ctx, dbconn, dbschema)
		if err != nil {
			log.Fatalf("Could not initialise DB: %s", err)
		}
	case "list":
		if len(os.Args) != 2 {
			log.Fatalf("Invalid usage. %s list", os.Args[0])
		}
		notes, err := cmd.ListAllNotes(ctx, dbconn)
		if err != nil {
			log.Fatalf("Could not list notes: %s", err)
		}
		NoteTableForTerminal(notes)
	case "create":
		if len(os.Args) != 4 {
			log.Fatalf("Invalid usage. %s create <title> <content>", os.Args[0])
		}
		note, err := cmd.CreateNote(ctx, dbconn, database.CreateNoteParams{
			Title: os.Args[2],
			Content: sql.NullString{
				String: os.Args[3],
				Valid:  true,
			},
		})
		if err != nil {
			log.Fatalf("Could not create note: %s", err)
		}
		log.Printf("Created note: %s\n", note.Title)
	case "select":
		if len(os.Args) != 3 {
			log.Fatalf("Invalid usage. %s select <id>", os.Args[0])
		}
		log.Fatal("Not implemented.")
	default:
		log.Fatal("Invalid command")
	}
}

func NoteTableForTerminal(notes []database.Note) {

	if len(notes) < 1 {
		fmt.Println("You have no notes.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tContent")
	for _, note := range notes {
		fmt.Fprintln(w, fmt.Sprintf("%d\t%s\t%s", note.ID, truncateString(note.Title, 30), truncateString(note.Content.String, 100)))
	}

	w.Flush()
}

// Credit Takahiro Kudo: https://dev.to/takakd/go-safe-truncate-string-9h0
func truncateString(str string, length int) string {
	if length <= 0 {
		return ""
	}

	// This code cannot support Japanese
	// orgLen := len(str)
	// if orgLen <= length {
	//     return str
	// }
	// return str[:length]

	// Support Japanese
	// Ref: Range loops https://blog.golang.org/strings
	truncated := ""
	count := 0
	for _, char := range str {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}
	return truncated
}
