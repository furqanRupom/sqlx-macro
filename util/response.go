package util

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, statusCode int, message string, errs ...error) {
    w.WriteHeader(statusCode)

    resp := map[string]string{
        "error": message,
    }

    if len(errs) > 0 && errs[0] != nil {
        resp["details"] = errs[0].Error()
    }

    json.NewEncoder(w).Encode(resp)
}

func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}