package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	db, err := sql.Open("sqlite", "file:test.db?cache=shared")
	if err != nil {
		panic(err)
	}
	a.db = db

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		panic(err)
	}
}

func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Insert() {
	_, err := a.db.Exec("INSERT INTO test (name) VALUES ('Hello, World!')")
	if err != nil {
		panic(err)
	}
}
