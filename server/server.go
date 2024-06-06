package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"apptodo.com/constant"

	"apptodo.com/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

// LOG CONFIG
func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Init Started!")
}

func main() {
	fmt.Println("user", constant.USER)
	fmt.Println("port", constant.PORT)
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", constant.USER, constant.PWD, constant.SERVER, constant.PORT, constant.BD)
	//Connect to database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	engine := html.New("../views", ".html")
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

	app.Static("/", "../public") // add this before starting the app

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

		fmt.Println(todo)

		todos = append(todos, todo)
	}

	return c.Render("index", fiber.Map{"Todos": todos})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := model.Todo{}

	err := c.BodyParser(&newTodo)

	if err != nil {
		log.Printf("An error ocurred: %v", err)
		return c.SendString(err.Error())
	}

	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT INTO todo (item) VALUES ($1)", newTodo.Item)

		if err != nil {
			log.Fatalf("An error occurred while executing query: %v", err)
		}

	}
	return c.Redirect("/")

}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	oldItem := c.Query("olditem")
	newItem := c.Query("newitem")
	id := c.Query("id")

	db.Exec("UPDATE todo SET item = $1 WHERE item = $2 AND id = $3", newItem, oldItem, id)

	return c.Redirect("/")

}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	id := c.Query("id")
	db.Exec("DELETE FROM todo WHERE id = $1", id)
	return c.SendString("Deleted!")

}
