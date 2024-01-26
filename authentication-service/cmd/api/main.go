package main

import (
	authentication_controller "authentication/internal/controller/authentication"
	http_server "authentication/internal/http"
	"authentication/internal/http/middleware"
	data "authentication/internal/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication Service")

	// Connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't Connect to Postgres")
	}

	//set up  config
	// app := &Config{
	// 	DB:     conn,
	// 	Models: data.New(conn),
	// }

	models := data.New(conn)

	srv := http_server.NewHTTPServer("debug")
	root := srv.Group("")

	root.Use(middleware.ErrorHandler())

	authentication_controller.AuthenticationController(root, models)

	srv.Run(fmt.Sprintf(":%s", webPort))
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}
		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
