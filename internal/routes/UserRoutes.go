package routes

import (
	"net/http"
	"project/internal/response"
	"project/internal/services"
	"project/internal/utils"

	"github.com/gorilla/mux"
)

// ApplyUserRoutes export
func ApplyUserRoutes(router *mux.Router) {
	router.HandleFunc(utils.BuildRouteURL(FetchProfileRoute), FetchProfileGet).Methods("GET")
}

// FetchProfileGet router
func FetchProfileGet(w http.ResponseWriter, r *http.Request) {
	deviceIDStr, err := utils.GetQueryParam(r, "deviceId")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	profile, err := services.FetchProfile(deviceIDStr)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, profile)
}
