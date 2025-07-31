package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/kai-gibson/workout_tracker/internal/app"
	"github.com/kai-gibson/workout_tracker/internal/config"
	model "github.com/kai-gibson/workout_tracker/internal/db"
	"github.com/kai-gibson/workout_tracker/internal/handler"
	"github.com/kai-gibson/workout_tracker/internal/logging"
)

func main() {
	if len(os.Args) > 1 {
		log.Printf("Loading env file: %s\n", os.Args[1])
		if err := godotenv.Load(os.Args[1]); err != nil {
			log.Fatalf("Error loading env file %v\n", err)
		}
	}

	config, err := config.NewFromEnv()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	db, err := sqlx.Connect("postgres", config.DbDsn())
	if err != nil {
		log.Fatalf("Error opening db: %v", err)
	}

	if !config.SkipMigrate() {
		err = model.MigrateDB(db.DB)
		if err != nil {
			log.Fatalf("Migration error: %v\n", err)
		}
	}

	app := &app.App{
		DB:     db,
		Config: &config,
		Logger: logging.NewLogger(config),
	}

	router := handler.NewRouter(app)
	app.Logger.Info("listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
