package course

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/notification/back-end/pkg/service/course"

	"github.com/notification/back-end/pkg/model"
)

func getAllCourse(service course.CourseServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, list_Course := service.GetAll(r.Context())
		err := json.NewEncoder(w).Encode(list_Course)
		if err != nil {
			ErroHttpMsgToConvertingResponseCourseListToJson.Write(w)
			return
		}
	})
}

func createCourse(service course.CourseServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course := &model.Courses{}
		name := r.URL.Query().Get("nome")

		if name == "" {
			http.Error(w, "o Nome do curso e obrigatório", http.StatusBadRequest)
			return
		}

		course.Name = name

		result, err := service.Create(r.Context(), course)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Error ou salvar turma"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}
}

func updateCourse(service course.CourseServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id_course := r.URL.Query().Get("id_curso")
		courseID, err := strconv.Atoi(id_course)
		if err != nil {
			http.Error(w, "Erro: id_turma não é um valor válido", http.StatusBadRequest)
			return
		}
		_, err = service.GetByID(r.Context(), courseID)
		if err != nil {
			http.Error(w, "curoso não encontrada", http.StatusNotFound)
			return
		}

		course := &model.Courses{}
		name := r.URL.Query().Get("nome")

		if name == "" {
			http.Error(w, "o horário e obrigatório", http.StatusBadRequest)
			return
		}

		course.Name = name
		course.ID = courseID
		_, err = service.Update(r.Context(), courseID, *&course)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Error ao atualizar Turma", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"MSG": "Success", "codigo": 1})
	}
}

func deleteCourse(service course.CourseServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_course := r.URL.Query().Get("id_turma")
		courseID, err := strconv.Atoi(id_course)
		if err != nil {
			http.Error(w, "Erro: id_turma não é um valor válido", http.StatusBadRequest)
			return
		}
		_, err = service.GetByID(r.Context(), courseID)
		if err != nil {
			http.Error(w, "Turma não encontrada", http.StatusNotFound)
			return
		}

		del, err := service.Delete(r.Context(), courseID)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Error ao atualizar Turma", http.StatusInternalServerError)
			return
		}

		if del {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"MSG": "Deletado com sucesso", "codigo": 1})

		}

	}
}

func getCourseById(service course.CourseServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id_course := r.URL.Query().Get("id_curso")
		courseID, err := strconv.Atoi(id_course)
		if err != nil {
			http.Error(w, "Erro: id_curso não é um valor válido", http.StatusBadRequest)
			return
		}

		Course, err := service.GetByID(r.Context(), courseID)
		if Course.ID == 0 {
			ErroHttpMsgCourseNotFound.Write(w)
			return
		}

		err = json.NewEncoder(w).Encode(Course)
		if err != nil {
			ErroHttpMsgToParseResponseCourseToJson.Write(w)
			return
		}
	})
}
