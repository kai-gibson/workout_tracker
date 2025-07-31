package handler

import (
	"net/http"

  "github.com/kai-gibson/workout_tracker/internal/app"

	"github.com/go-chi/chi/v5"
)

func NewRouter(app *app.App) http.Handler {
	router := chi.NewRouter()

	router.Get("/user", app.ListUsers)
  router.Post("/user", app.CreateUser)

	return router
}
