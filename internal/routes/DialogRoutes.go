package routes

import (
	"encoding/json"
	"net/http"

	"project/internal/logger"
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
		response.Make(w, response.StatusFail, err)
		return
	}

	dialogs, err := services.FetchDialogs(profileIDStr)
	if err != nil {
		response.Make(w, response.StatusFail, err)
		return
	}

	response.Make(w, response.StatusSuccess, dialogs)
}

// GetLoadDialog load dialog
func GetLoadDialog(w http.ResponseWriter, r *http.Request) {
	profileIDStr, err := utils.GetQueryParam(r, "profileId")
	if err != nil {
		response.Make(w, response.StatusFail, err)
		return
	}

	userIDStr, err := utils.GetQueryParam(r, "userId")
	if err != nil {
		response.Make(w, response.StatusFail, err)
		return
	}

	dialog, err := services.LoadDialog(profileIDStr, userIDStr)
	if err != nil {
		response.Make(w, response.StatusFail, err)
		return
	}

	response.Make(w, response.StatusSuccess, dialog)
}

// PostCreateDialog create new dialog
func PostCreateDialog(w http.ResponseWriter, r *http.Request) {
	var b struct {
		profileIDStr string
		userIDStr    string
	}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.Make(w, response.StatusFail, err)
		return
	}

	dialog, err := services.CreateDialog(b.profileIDStr, b.userIDStr)
	if err != nil {
		response.Make(w, response.StatusFail, err)
		return
	}

	response.Make(w, response.StatusSuccess, dialog)
}
