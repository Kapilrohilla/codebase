package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kapilrohilla/codebase/internal/storage"
	"github.com/kapilrohilla/codebase/internal/types"
	response "github.com/kapilrohilla/codebase/internal/utils"
)

func New(db storage.Storage) http.HandlerFunc {

	handlePostMethod := func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		var err error = json.NewDecoder(r.Body).Decode(&student)

		if err != nil && errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// handle body
		slog.Info("creating a student")

		lastId, err := db.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			response.WriteJson(w, 500, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]interface{}{"success": "Ok", "id": lastId})
	}

	handleGetMethod := func(w http.ResponseWriter, _ *http.Request) {
		result, err := db.GetStudent(0, 10)
		if err != nil {
			fmt.Println(err)
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		var out map[string]interface{} = map[string]interface{}{"isValid": true, "data": result}
		response.WriteJson(w, http.StatusOK, out)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var method string = r.Method

		switch method {
		case "GET":
			handleGetMethod(w, r)
		case "POST":
			handlePostMethod(w, r)
		default:
			response.WriteJson(w, 404, response.GeneralError(fmt.Errorf("invalid request method")))
			return
		}
	}
}

func NewById(db storage.Storage) http.HandlerFunc {
	handleGetById := func(w http.ResponseWriter, r *http.Request) {
		var id string = r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		data, err := db.GetStudentById(intId)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		var out map[string]interface{} = map[string]interface{}{"isValid": true, "data": data}
		response.WriteJson(w, http.StatusOK, out)
	}
	handeUpdateById := func(w http.ResponseWriter, r *http.Request) {
		var id string = r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		data, err := db.GetStudentById(intId)

		if data.Id == 0 {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("student not found")))
		}
		var studentUpdatePayload types.Student

		json.NewDecoder(r.Body).Decode(&studentUpdatePayload)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		var out map[string]interface{} = map[string]interface{}{"isValid": true, "data": data}
		response.WriteJson(w, http.StatusOK, out)
	}
	return func(w http.ResponseWriter, r *http.Request) {

		var method string = r.Method

		switch method {
		case "GET":
			handleGetById(w, r)
		case "PUT":
			handeUpdateById(w, r)
		default:
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("route not found")))
		}

	}
}
