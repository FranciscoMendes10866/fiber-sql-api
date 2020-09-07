package handler

import (
	"database/sql"
	"fmt"
	"log"

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

func FindAllUsers(c *fiber.Ctx) {
	type UsersData struct {
		ID       int
		Username string
		Email    string
		Password string
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
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users []UsersData
	for rows.Next() {
		var u UsersData
		// queries the properties I want
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		if err != nil {
			panic(err.Error())
		}
		// appends those properties to the users var
		users = append(users, u)
		// response
		c.JSON(users)
	}
}
