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

// ApplyUserRoutes export
func ApplyUserRoutes(router *mux.Router) {
	router.HandleFunc(utils.BuildRouteURL(FetchProfileRoute), FetchProfileGet).Methods(http.MethodGet)
	router.HandleFunc(utils.BuildRouteURL(UpdateProfileRoute, "id"), UpdateProfilePut).Methods(http.MethodPut)
	router.HandleFunc(utils.BuildRouteURL(FetchProfileCountsRoute), FetchProfileCountsGet).Methods(http.MethodGet)
	router.HandleFunc(utils.BuildRouteURL(FetchPeopleUsersRoute), FetchPeopleUsers).Methods(http.MethodGet)
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

// UpdateProfilePut create new user
func UpdateProfilePut(w http.ResponseWriter, r *http.Request) {
	profileIDStr, err := utils.GetQueryParam(r, "id")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	var b models.User

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	updatedProfile, err := services.UpdateProfile(profileIDStr, &b)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, updatedProfile)
}

// FetchProfileCountsGet fetch counts and profile
func FetchProfileCountsGet(w http.ResponseWriter, r *http.Request) {
	profileIDStr, err := utils.GetQueryParam(r, "profileId")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	profileCounts, err := services.FetchProfileCounts(profileIDStr)
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, profileCounts)
}

// FetchPeopleUsers fetch people users
func FetchPeopleUsers(w http.ResponseWriter, r *http.Request) {
	queryParams, err := utils.GetQueryParams(r, "profileId", "filterOptions", "usersOffset")
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	peopleUsers, err := services.FetchPeopleUsers(queryParams["profileId"], queryParams["filterOptions"], queryParams["usersOffset"])
	if err != nil {
		response.MakeResponseObject(w, response.StatusFail, err)
		return
	}

	response.MakeResponseObject(w, response.StatusSuccess, peopleUsers)
}
