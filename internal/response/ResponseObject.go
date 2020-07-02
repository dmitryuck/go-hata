package response

import (
	"encoding/json"
	"net/http"
)

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
)

type ResponseObject struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}

// Make ResponseObject
func (r ResponseObject) Make(w http.ResponseWriter, status string, payload interface{}) {
	responseObject := &ResponseObject{
		Status:  status,
		Payload: payload,
	}

	jsonResponse, err := json.MarshalIndent(responseObject, "", " ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonResponse)))
}
