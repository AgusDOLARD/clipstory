package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	storage "github.com/AgusDOLARD/clipstory/sqlc"
)

func main() {
	clearFlag := flag.Bool("d", false, "clear clips")
	flag.Parse()

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "/tmp/clipstory.db")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	q := storage.New(db)

	if *clearFlag {
		err := q.Clear(ctx)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		os.Exit(0)
	}

	clips, err := q.GetAll(ctx)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	for _, e := range clips {
		fmt.Println(e.Value)
	}
}
