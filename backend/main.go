package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// routes for the app
func routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi ya")
	})
}

// newDB creates a database connection
func newDB() (*sql.DB, error) {
	// Environment variables for port, user, password, and dbname
	port := os.Getenv("ENV_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dbConnection := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, dbname)
	db, err := sql.Open("postgres", dbConnection)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := newDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()
	routes(app)
	log.Fatal(app.Listen(":3000"))
}
