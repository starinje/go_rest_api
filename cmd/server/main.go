package main

import (
	"fmt"

	"go_rest_api/internal/comment"

	"go_rest_api/internal/db"

	transportHttp "go_rest_api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil

}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
