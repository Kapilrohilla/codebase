package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/kapilrohilla/codebase/internal/types"
	response "github.com/kapilrohilla/codebase/internal/utils"
)

func New() http.HandlerFunc {

	handlePostMethod := func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		var err error = json.NewDecoder(r.Body).Decode(&student)

		if err != nil && errors.Is(err, io.EOF) {
			// throw error
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// handle body
		slog.Info("creating a student")
		response.WriteJson(w, http.StatusCreated, map[string]interface{}{"success": "Ok"})
	}

	handleGetMethod := func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Welcome to GO-based student-service\nWritten by Kapil Rohilla"))
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
