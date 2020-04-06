package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App : for holding our application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize : for connection to database
func (a *App) Initialize(user, password, dbname, sslmode string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run : for simply start the application
func (a *App) Run(addr string) {

}
