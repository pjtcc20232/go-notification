package course

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/notification/back-end/pkg/service/course"
)

func RegisterCourseAPIHandlers(r chi.Router, service course.CourseServiceInterface) {
	r.Route("/api/v1/course", func(r chi.Router) {
		r.Post("/add", createCourse(service))
		r.Put("/update/{id_course}/{nome}", updateCourse(service))
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			handler := getAllCourse(service)
			handler.ServeHTTP(w, r)
		})
	})

}
