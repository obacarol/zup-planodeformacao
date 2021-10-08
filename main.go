package main

import (
	"fmt"
	"log"
	"net/http"

	"planodeformacao-upgrade/connectionDB"
	"planodeformacao-upgrade/router"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db := connectionDB.DbConn()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://C:/Users/caroline.silva/go/src/github.com/caroline.silva-zup/planodeformacao-upgrade/connectionDB/migration",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Steps(-2)
	m.Steps(2)

	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
