package class

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notification/back-end/pkg/service/class"
)

func RegisterClassAPIHandlers(r chi.Router, service class.ClassServiceInterface) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/class", createClass(service))
		r.Put("/class/{id}", updateClass(service))
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			handler := getAllClass(service)
			handler.ServeHTTP(w, r)
		})
	})

}
