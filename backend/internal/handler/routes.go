package handler

import (
	"net/http"

  "github.com/kai-gibson/workout_tracker/internal/app"

	"github.com/go-chi/chi/v5"
)

func NewRouter(app *app.App) http.Handler {
	router := chi.NewRouter()

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message":"Hi There:)"}`))
	})

	router.Get("/users", app.ListUsers)

	return router
}
