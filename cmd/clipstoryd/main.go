package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.design/x/clipboard"

	storage "github.com/AgusDOLARD/clipstory/sqlc"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "/tmp/clipstory.db")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	err = storage.Migrate(ctx, db)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	q := storage.New(db)

	ch := clipboard.Watch(ctx, clipboard.FmtText)
	for e := range ch {
		err := q.Create(ctx, string(e))
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}

}
