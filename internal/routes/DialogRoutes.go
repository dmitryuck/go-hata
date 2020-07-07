package routes

import (
	"encoding/json"
	"net/http"

	"project/internal/models"
	"project/internal/response"
	"project/internal/services"
	"project/internal/utils"

	"github.com/gorilla/mux"
)

// ApplyDialogRoutes for methods
func ApplyDialogRoutes(router *mux.Router) {
	router.HandleFunc(utils.BuildRouteURL(FetchDialogsRoute), FetchDialogsGet).Methods("GET")
	router.HandleFunc(utils.BuildRouteURL(LoadDialogRoute), LoadDialogGet).Methods("GET")
	router.HandleFunc(utils.BuildRouteURL(CreateDialogRoute), CreateDialogPost).Methods("POST")
	router.HandleFunc(utils.BuildRouteURL(SendMessageRoute), SendMessagePost).Methods("POST")
	router.HandleFunc(utils.BuildRouteURL(DeleteDialogRoute), DeleteDialogPost).Methods("POST")
}

// FetchDialogsGet router
func FetchDialogsGet(w http.ResponseWriter, r *http.Request) {
	profileIDStr, err := utils.GetQueryParam(r, "profileId")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	dialogs, err := services.FetchDialogs(profileIDStr)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, dialogs)
}

// LoadDialogGet load dialog
func LoadDialogGet(w http.ResponseWriter, r *http.Request) {
	profileIDStr, err := utils.GetQueryParam(r, "profileId")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	userIDStr, err := utils.GetQueryParam(r, "userId")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	dialog, err := services.LoadDialog(profileIDStr, userIDStr)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, dialog)
}

// CreateDialogPost create new dialog
func CreateDialogPost(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		userIDStr    string
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	dialog, err := services.CreateDialog(b.profileIDStr, b.userIDStr)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, dialog)
}

// SendMessagePost send message
func SendMessagePost(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		dialogIDStr  string
		text         string
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	updatedDialog, err := services.SendMessage(b.profileIDStr, b.dialogIDStr, b.text)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, updatedDialog)
}

// DeleteDialogPost delete dialog
func DeleteDialogPost(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		dialogIDStr  string
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	deleteResult, err := services.DeleteDialog(b.profileIDStr, b.dialogIDStr)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, deleteResult)
}

// UpdateDialogPut update dialog
func UpdateDialogPut(w http.ResponseWriter, r *http.Request) {
	dialogIDStr, err := utils.GetQueryParam(r, "id")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	var b models.Dialog

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	updatedDialog, err := services.UpdateDialog(dialogIDStr, &b)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, updatedDialog)
}
