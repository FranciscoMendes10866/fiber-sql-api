package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber"
)

func CreateUser(c *fiber.Ctx) {
	type InputData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// input data
	input := new(InputData)
	c.BodyParser(input)
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
	// inserts the user
	prep, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	prep.Exec(input.Username, input.Email, input.Password)
	// response
	c.SendStatus(201)
}

func LogInUser(c *fiber.Ctx) {
	type InputData struct {
		Email string `json:"email"`
	}
	type UserData struct {
		ID       int
		Username string
		Email    string
		Password string
	}
	// input data
	input := new(InputData)
	c.BodyParser(input)
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

	rows, err := db.Query("SELECT * FROM users WHERE email = ?", input.Email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users UserData
	for rows.Next() {
		var u UserData
		// queries the properties I want
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		if err != nil {
			panic(err.Error())
		}
		// appends those properties to the users var
		users = u
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = users.ID
	// Generate encoded token and send it as response.
	loginToken, err := token.SignedString([]byte("FIBERSQL"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}
	// Response
	c.JSON(fiber.Map{
		"token": loginToken,
	})
}
