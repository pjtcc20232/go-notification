package class

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notification/back-end/pkg/service/class"
)

func RegisterClassAPIHandlers(r chi.Router, service class.ClassServiceInterface) {
	r.Route("/api/v1/class", func(r chi.Router) {
		r.Post("/add", createClass(service))
		r.Put("/update/{id}", updateClass(service))
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			handler := getAllClass(service)
			handler.ServeHTTP(w, r)
		})
	})

}
