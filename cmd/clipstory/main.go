package main

import (
	"fmt"
	"os"

	"github.com/AgusDOLARD/clipstory/data"
	database "github.com/AgusDOLARD/clipstory/db"
)

func main() {
	err := database.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	entries, err := data.GetEntries()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, e := range entries {
		fmt.Println(e.Value)
	}
}
