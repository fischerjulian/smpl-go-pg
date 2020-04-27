package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	_ "github.com/lib/pq"
)

func main() {

	postgresUsername := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")

	connStr := "user=" + postgresUsername + " dbname=postgres sslmode=disable password=" + postgresPassword + " host=" + postgresHost
	db, err := sql.Open("postgres", connStr)
	company := ""
	if err != nil {
		log.Fatal(err)
	}

	app := iris.New()
	app.Logger().SetLevel("debug")

	// recover panics
	app.Use(recover.New())
	app.Use(logger.New())

	// Method: GET
	// Resource http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		rows, err := db.Query("SELECT NAME FROM COMPANY")
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err = rows.Scan(&company)
			if err != nil {
				log.Fatal(err)
			}
		}

		ctx.HTML("<h1>Simple PostgreSQL Web App</h1><h2>Company: " + company + "<h2>")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
