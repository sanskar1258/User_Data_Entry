package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"dataentry/router"
	"dataentry/utility"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found...")
	}

	db_user := os.Getenv("DB_USER")
	print(db_user)
	if db_user == "" {
		log.Fatal("Missing env value DB_USER")
	}

	db_pass := os.Getenv("DB_PASSWORD")
	print(db_pass)
	if db_pass == "" {
		log.Fatal("Missing env value DB_PASSWORD")
	}

	db_name := os.Getenv("DB_NAME")
	print(db_name)
	if db_name == "" {
		log.Fatal("Missing env value DB_NAME")
	}

	// open connection
	utility.Db, err = sql.Open("mysql", db_user+":"+db_pass+"@tcp(127.0.0.1:3306)/"+db_name)

	if err != nil {
		panic(err)
	}

	utility.Db.SetConnMaxLifetime(time.Minute * 3)

}

func main() {
	fmt.Println("Build a simple web server with golang")

	web_port := os.Getenv("WEB_PORT")
	if web_port == "" {
		log.Fatal("Misiing env value WEB_PORT")
	}

	routerHandler := router.RouteHandler()

	fmt.Printf("Starting server at port: %s\n", web_port)
	if err := http.ListenAndServe(":"+web_port, routerHandler); err != nil {
		log.Fatal(err)
	}

}
