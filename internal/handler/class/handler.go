package class

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/notification/back-end/pkg/model"
	"github.com/notification/back-end/pkg/service/class"
)

func getAllClass(service class.ClassServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, list_class := service.GetAll(r.Context())
		err := json.NewEncoder(w).Encode(list_class)
		if err != nil {
			ErroHttpMsgToConvertingResponseClassListToJson.Write(w)
			return
		}
	})
}

func createClass(service class.ClassServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		class := &model.Class{}
		schedule := r.URL.Query().Get("horario")
		id_course := r.URL.Query().Get("id_turma")
		courseID, err := strconv.Atoi(id_course)
		if err != nil {
			http.Error(w, "Erro: id_turma não é um valor válido", http.StatusBadRequest)
			return
		}

		if schedule == "" {
			http.Error(w, "o horário e obrigatório", http.StatusBadRequest)
			return
		}

		class.Schedules = schedule
		class.Tbl_course_id = courseID

		result, err := service.Create(r.Context(), class)
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

func updateClass(service class.ClassServiceInterface) http.HandlerFunc {
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

		class := &model.Class{}
		schedule := r.URL.Query().Get("horario")

		if schedule == "" {
			http.Error(w, "o horário e obrigatório", http.StatusBadRequest)
			return
		}

		class.Schedules = schedule
		class.Tbl_course_id = courseID
		_, err = service.Update(r.Context(), courseID, *&class)
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

func deleteClass(service class.ClassServiceInterface) http.HandlerFunc {
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

func getClassById(service class.ClassServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id_course := r.URL.Query().Get("id_turma")
		courseID, err := strconv.Atoi(id_course)
		if err != nil {
			http.Error(w, "Erro: id_turma não é um valor válido", http.StatusBadRequest)
			return
		}

		class, err := service.GetByID(r.Context(), courseID)
		if class.ID == 0 {
			ErroHttpMsgClassNotFound.Write(w)
			return
		}

		err = json.NewEncoder(w).Encode(class)
		if err != nil {
			ErroHttpMsgToParseResponseClassToJson.Write(w)
			return
		}
	})
}
