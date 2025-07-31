package app

import (
  "log/slog"
  "net/http"
  "encoding/json"

	dbmodel "github.com/kai-gibson/workout_tracker/internal/db"
  "github.com/kai-gibson/workout_tracker/internal/config"
  "github.com/jmoiron/sqlx"
)

// context for all endpoints
type App struct {
  DB *sqlx.DB
  Config *config.Config
  Logger *slog.Logger
}

func (app *App) ListUsers(w http.ResponseWriter, r *http.Request) {
  users := []dbmodel.User{}
  err := app.DB.Select(&users, "SELECT * FROM tbl_users")
  if err != nil {
    app.Logger.Warn("Failed to list users", "error", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  app.Logger.Info("Users response", "users", users)
  json.NewEncoder(w).Encode(users)
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
  type UserBody struct {
    Email string `json:"email"`
    Username string `json:"username"`
    HashedPwd string `json:"hashed_password"`
  }
  
  var body UserBody
  decoder := json.NewDecoder(r.Body)
  decoder.DisallowUnknownFields()

  err := decoder.Decode(&body)
  if err != nil {
    http.Error(w, "Body decode error", http.StatusInternalServerError)
    return
  }

  json.NewEncoder(w).Encode(body)
}
