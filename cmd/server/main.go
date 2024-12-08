package main

import (
	"context"
	"fmt"

	"github.com/TutorialEdge/go-rest-api-course/internal/db"
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
