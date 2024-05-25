package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//IMPORT LOCAL MODULE
	model "apptodo/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://postgres:admin@localhost:5432/testdb?sslmode=disable"
	//Connect to database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("/", "./public") // add this before starting the app

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	todos := make([]*model.Todo, 0)

	rows, err := db.Query("SELECT id, item FROM todo")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON("Erro na execução da consulta no BD!")
	}

	for rows.Next() {
		todo := new(model.Todo)

		rows.Scan(&todo.Id, &todo.Item)

		fmt.Println(model.Todo.PrintInfo(*todo))

		todos = append(todos, todo)
	}

	return c.Render("index", fiber.Map{"Todos": todos})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")

}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")

}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello ")

}
