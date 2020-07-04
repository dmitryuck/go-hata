package routes

import (
	"encoding/json"
	"net/http"

	"project/internal/logger"
	"project/internal/models"
	"project/internal/response"
	"project/internal/services"
	"project/internal/utils"

	"github.com/gorilla/mux"
)

// ApplyDialogRoutes for methods
func ApplyDialogRoutes(router *mux.Router) {
	//router.HandleFunc(utils.BuildRouteURL(""), GetTest).Methods("GET")
	router.HandleFunc(utils.BuildRouteURL(FetchDialogsRoute), GetFetchDialogs).Methods("GET")
	router.HandleFunc(utils.BuildRouteURL(LoadDialogRoute), GetLoadDialog).Methods("GET")
	router.HandleFunc(utils.BuildRouteURL(CreateDialogRoute), PostCreateDialog).Methods("POST")
	router.HandleFunc(utils.BuildRouteURL(SendMessageRoute), PostSendMessage).Methods("POST")
	router.HandleFunc(utils.BuildRouteURL(DeleteDialogRoute), PostDeleteDialog).Methods("POST")
}

// GetTest test route
func GetTest(w http.ResponseWriter, r *http.Request) {
	result := utils.BuildRouteURL(FetchDialogsRoute, "id")
	logger.Instance.LogInfo(result)
}

// GetFetchDialogs router
func GetFetchDialogs(w http.ResponseWriter, r *http.Request) {
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

// GetLoadDialog load dialog
func GetLoadDialog(w http.ResponseWriter, r *http.Request) {
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

// PostCreateDialog create new dialog
func PostCreateDialog(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		userIDStr    string
	}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
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

// PostSendMessage send message
func PostSendMessage(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		dialogIDStr  string
		text         string
	}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
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

// PostDeleteDialog delete dialog
func PostDeleteDialog(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		dialogIDStr  string
	}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
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

// PutUpdateDialog update dialog
func PutUpdateDialog(w http.ResponseWriter, r *http.Request) {
	dialogIDStr, err := utils.GetQueryParam(r, "id")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	var b models.Dialog

	decodeErr := json.NewDecoder(r.Body).Decode(&b)
	if decodeErr != nil {
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
