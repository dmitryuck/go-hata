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

// MakeResponseObject make ResponseObject
func MakeResponseObject(w http.ResponseWriter, status string, payload interface{}) {
	responseObject := &ResponseObject{
		Status:  status,
		Payload: payload,
	}

	if status == "fail" {
		responseObject.Payload = payload.(error).Error()
	}

	jsonResponse, err := json.MarshalIndent(responseObject, "", " ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonResponse)))
}
