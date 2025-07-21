package main

import (
	_ "database/sql"
	"log"
	"net/http"

	dbmodel "github.com/kai-gibson/workout_tracker/internal/db"
	"github.com/jmoiron/sqlx"
	"github.com/kai-gibson/workout_tracker/internal/handler"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	db, err := sqlx.Connect(
		"postgres",
		"user=kai dbname=workout_tracker sslmode=disable password=secret",
	)
	if err != nil {
		log.Fatalf("Error opening db: %v", err)
	}

  err = dbmodel.MigrateDB(db.DB)
  if err != nil {
    log.Fatalf("Migration error: %v\n", err)
  }

	user := dbmodel.User{}
	err = db.Get(&user, "SELECT * FROM tbl_users")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello there")

	router := handler.NewRouter()
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
