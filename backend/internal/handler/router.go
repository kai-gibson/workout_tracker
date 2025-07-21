package handler

import (
  "net/http"
  "github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
  router := chi.NewRouter()

  router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("content-type", "application/json")
    w.Write([]byte(`{"message":"Hi There:)"}`))
  })

  router.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
    userID := chi.URLParam(r, "id")
    w.Header().Add("content-type", "application/json")
    w.Write([]byte(`{"message":"User ID: ` + userID + ` OK?"}`))
  })

  return router
}
