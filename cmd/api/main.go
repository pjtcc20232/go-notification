package main

import (
	"log"
	"net/http"

	"github.com/notification/back-end/internal/config"
	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/internal/handler/class"
	"github.com/notification/back-end/internal/handler/course"
	"github.com/notification/back-end/pkg/adapter/pgsql"

	"github.com/notification/back-end/pkg/server"

	class_service "github.com/notification/back-end/pkg/service/class"
	course_service "github.com/notification/back-end/pkg/service/course"

	"github.com/go-chi/chi/v5"
)

var (
	VERSION = "0.1.0-dev"
	COMMIT  = "ABCDEFG-dev"
)

func main() {

	logger.Info("start Notifaction application")
	conf := config.NewConfig()

	db_pool := pgsql.New(conf)
	class_service := class_service.NewClassService(db_pool)
	course_service := course_service.NewCourseService(db_pool)
	r := chi.NewRouter()

	r.Get("/", healthcheck)
	class.RegisterClassAPIHandlers(r, class_service)
	course.RegisterCourseAPIHandlers(r, course_service)
	srv := server.NewHTTPServer(r, conf)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Printf("Server Run on [Port: %s], [Mode: %s], [Version: %s], [Commit: %s]", conf.PORT, conf.Mode, VERSION, COMMIT)

	select {}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"MSG": "Server Ok", "codigo": 200}`))
}
