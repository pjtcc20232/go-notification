package course

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/service/course"

	"github.com/notification/back-end/pkg/model"
)

func getAllCourse(service course.CourseServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		courseList, err := service.GetAll(r.Context())

		if err != nil {
			// Lida com o erro, por exemplo, loga-o ou retorna uma resposta de erro.
			//logger.Error("Erro ao obter a lista de cursos: " + err.Error())
			http.Error(w, "Erro ao obter a lista de cursos", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(courseList)
		if err != nil {
			// Lida com o erro de codificação JSON, se ocorrer.
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

		_, err := service.Create(r.Context(), course)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Error ou salvar Curso"+err.Error(), http.StatusInternalServerError)
			return
		}

		type Response struct {
			Message string `json:"message"`
		}

		// Crie uma instância da estrutura com a mensagem desejada.
		msg := Response{
			Message: "Dados gravados com sucesso",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(msg)
	}
}

func updateCourse(service course.CourseServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id_course := chi.URLParam(r, "id_course")
		logger.Info("PEGANDO O PARAMENTRO")
		logger.Info(id_course)
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
		nome := chi.URLParam(r, "nome")
		logger.Info("PEGANDO O NOME")
		logger.Info(nome)
		if nome == "" {
			http.Error(w, "o Nome do curso e obrigatório", http.StatusBadRequest)
			return
		}

		course.Name = nome
		course.ID = courseID
		_, err = service.Update(r.Context(), courseID, *&course)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Error ao atualizar cruso", http.StatusInternalServerError)
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
			http.Error(w, "Erro: id_curso não é um valor válido", http.StatusBadRequest)
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

		course, err := service.GetByID(r.Context(), courseID)
		if course.ID == 0 {
			ErroHttpMsgCourseNotFound.Write(w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(course)
		if err != nil {
			ErroHttpMsgToParseResponseCourseToJson.Write(w)
			return
		}
	})
}
