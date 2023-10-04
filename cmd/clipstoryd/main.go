package main

import (
	"context"
	"fmt"
	"os"

	"golang.design/x/clipboard"

	"github.com/AgusDOLARD/clipstory/data"
	database "github.com/AgusDOLARD/clipstory/db"
)

func main() {
	ctx := context.TODO()

	err := database.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = database.Db.AutoMigrate(&data.Entry{})
	if err != nil {
		fmt.Println(err)
	}

	ch := clipboard.Watch(ctx, clipboard.FmtText)
	for e := range ch {
		entry := &data.Entry{Value: string(e)}
		err := entry.Save()
		if err != nil {
			fmt.Println(err)
		}
	}

}
