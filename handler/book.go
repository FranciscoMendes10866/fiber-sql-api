package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func CreateBook(c *fiber.Ctx) {
	type InputData struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	// input data
	input := new(InputData)
	c.BodyParser(input)
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(float64)
	// converts the tokenID from float64 (1.0000) to int (1)
	var IDtoInt int = int(tokenID)
	// connects to db
	db, err := sql.Open("mysql", "root:root@tcp(localhost:7788)/fiber")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the MySQL database! 😭")
	// defer will close the connection when the main function has finished
	defer db.Close()
	// inserts the book
	prep, err := db.Prepare("INSERT INTO books (title, author, user_id) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	prep.Exec(input.Title, input.Author, IDtoInt)
	// response
	c.SendStatus(201)
}

func GetAllBooks(c *fiber.Ctx) {
	type UsersData struct {
		ID     int
		Title  string
		Author string
		UserID int
	}
	// connects to db
	db, err := sql.Open("mysql", "root:root@tcp(localhost:7788)/fiber")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the MySQL database! 😭")
	// defer will close the connection when the main function has finished
	defer db.Close()
	// Queries and gets all users
	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var books []UsersData
	for rows.Next() {
		var b UsersData
		// queries the properties I want
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.UserID)
		if err != nil {
			panic(err.Error())
		}
		// appends those properties to the users var
		books = append(books, b)
		// response
		c.JSON(books)
	}
}

func DeleteBook(c *fiber.Ctx) {
	BookID := c.Params("id")
	// connects to db
	db, err := sql.Open("mysql", "root:root@tcp(localhost:7788)/fiber")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the MySQL database! 😭")
	// defer will close the connection when the main function has finished
	defer db.Close()
	// deletes the book
	delete, err := db.Exec(`DELETE FROM books WHERE id = ?`, BookID)
	if err != nil || delete == nil {
		panic(err.Error())
	}
	c.SendStatus(200)
}
