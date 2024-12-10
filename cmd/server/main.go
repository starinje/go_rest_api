package main

import (
	"context"
	"fmt"

	"go_rest_api/internal/comment"

	"go_rest_api/internal/db"
)

// Run - is going to be responsible for
// the initialization and startup of our
// go application
func Run() error {
	fmt.Println("Starting up our go application...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to connect to the database")
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "ff1d4530-959b-4a2d-a12f-47c095350e1b",
			Slug:   "my-first-comment",
			Author: "Elliott",
			Body:   "Hello, World",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"ff1d4530-878b-4a2d-a12f-47c095350e1b",
	))

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("Successfully connected and pinged database!")

	return nil

}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
