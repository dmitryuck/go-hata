package routes

import (
	"net/http"
	"project/internal/logger"
	"project/internal/utils"

	"github.com/gorilla/mux"
)

// ApplyRoutes for root
func ApplyRoutes(router *mux.Router) {
	//router.HandleFunc(utils.BuildRouteURL(""), GetTest).Methods("GET")

	ApplyDialogRoutes(router)
	ApplyUserRoutes(router)
}

// GetTest test route
func GetTest(w http.ResponseWriter, r *http.Request) {
	result := utils.BuildRouteURL(FetchDialogsRoute, "id")
	logger.Instance.LogInfo(result)
}
