package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

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
		// w.Write([]byte("Welcome to GO-based student-service\nWritten by Kapil Rohilla"))
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
