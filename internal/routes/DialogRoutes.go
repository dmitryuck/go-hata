package routes

import (
	"net/http"

	"project/internal/logger"
	"project/internal/utils"

	"github.com/gorilla/mux"
)

// ApplyDialogRoutes for methods
func ApplyDialogRoutes(router *mux.Router) {
	r := utils.BuildRouteURL(FetchDialogsRoute, "id")
	logger.Instance.LogInfo(r)

	router.HandleFunc(utils.BuildRouteURL(FetchDialogsRoute), GetFetchDialogs).Methods("GET")
}

// GetFetchDialogs router
func GetFetchDialogs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("Hello world")))
}
