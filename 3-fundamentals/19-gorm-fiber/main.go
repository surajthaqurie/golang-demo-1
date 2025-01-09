package main

import (
	"fmt"
	"gorm-fiber/book"
	"gorm-fiber/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

// homePage is the handler for the root route.
func homePage(c *fiber.Ctx) error {
	c.Send([]byte("Hello World!"))
	return nil

}

// setupRoutes defines the routes for the book API.
func setupRoutes(app *fiber.App) {

	//@Define the root route and assign the homePage handler
	app.Get("/", homePage)

	//@Define the routes for the book API
	app.Get("/api/v1/books", book.GetBooks)         //Get all books
	app.Get("/api/v1/book/:id", book.GetBook)       //Get a specific book by ID
	app.Post("/api/v1/book", book.AddBook)          //Add a new book
	app.Delete("/api/v1/book/:id", book.DeleteBook) //Delete a book by ID
}

// initDatabase initializes the MySQL database connection.
func initDatabase() {

	//@Read environment variables from the .env file
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//@Construct the MySQL connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	//@Open the database connection using GORM
	var err error
	database.DBconn, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected successfully opened")

	//@Automatically migrate the schema
	database.DBconn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

// main starts the server and initializes all necessary components.
func main() {
	fmt.Println("GORM with fiber framework")

	//@Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//@Create a new Fiber application
	app := fiber.New()

	//@Initialize the database
	initDatabase()

	//@Setup the API routes
	setupRoutes(app)

	//@Start the Fiber server on port 8000
	app.Listen(":8000")
}
