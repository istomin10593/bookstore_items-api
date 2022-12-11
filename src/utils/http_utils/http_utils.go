package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/istomin10593/bookstore_utils-go/rest_errors"
)

func ResponseJson(w http.ResponseWriter, statuscode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err rest_errors.RestErr) {
	ResponseJson(w, err.Status(), err)
}
