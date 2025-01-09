package book

import (
	"gorm-fiber/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Book defines the structure of a book in the system.
type Book struct {
	gorm.Model        // This automatically includes CreatedAt, UpdatedAt, and DeletedAt fields
	Title      string `json:"title"`
	Author     string `json:"author"`
	Rating     int    `json:"rating"`
}

// GetBooks fetches all books from the database and returns them in JSON format.
func GetBooks(c *fiber.Ctx) error {
	// Initialize an empty slice of books
	var books []Book

	// Fetch all books from the database
	if err := database.DBconn.Find(&books).Error; err != nil {
		c.Status(500).Send([]byte("Error fetching books"))
		return err
	}

	// Send the list of books as a JSON response
	c.JSON(books)
	return nil
}

// GetBook fetches a single book by its ID from the database and returns it in JSON format.
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id") // Get the book ID from the URL parameters

	// Initialize a book object
	var book Book

	// Fetch the book by its ID from the database
	if err := database.DBconn.First(&book, id).Error; err != nil {
		c.Status(404).Send([]byte("No book found with the given ID"))
		return err
	}

	// Send the book as a JSON response
	c.JSON(book)
	return nil
}

// AddBook creates a new book in the database using data from the request body.
func AddBook(c *fiber.Ctx) error {
	db := database.DBconn

	// Initialize a book object
	book := new(Book) // var book Book

	// Parse the JSON body and populate the book struct
	if err := c.BodyParser(&book); err != nil {
		c.Status(400).Send([]byte("Invalid input"))
		return err
	}

	// Create the new book in the database
	if err := db.Create(&book).Error; err != nil {
		c.Status(500).Send([]byte("Error creating book"))
		return err
	}

	// Send the newly created book as a JSON response
	c.JSON(book)
	return nil
}

// DeleteBook deletes a book by its ID from the database.
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id") // Get the book ID from the URL parameters

	// Initialize a book object
	var book Book

	// Find the book by its ID
	if err := database.DBconn.First(&book, id).Error; err != nil {
		c.Status(404).Send([]byte("No book found with the given ID"))
		return err
	}

	// Delete the book from the database
	if err := database.DBconn.Delete(&book).Error; err != nil {
		c.Status(500).Send([]byte("Error deleting book"))
		return err
	}

	// Send a success message
	c.Send([]byte("Book successfully deleted"))
	return nil
}
