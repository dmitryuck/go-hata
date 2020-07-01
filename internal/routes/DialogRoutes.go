package routes

import (
	"context"
	"net/http"

	"project/internal/db"
	"project/internal/logger"
	"project/internal/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// ApplyDialogRoutes for methods
func ApplyDialogRoutes(router *mux.Router) {
	//r := utils.BuildRouteURL(FetchDialogsRoute, "id")
	//logger.Instance.LogInfo(r)

	router.HandleFunc(utils.BuildRouteURL(FetchDialogsRoute), GetUsersCount).Methods("GET")
}

// GetFetchDialogs router
func GetFetchDialogs(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(string("Hello world")))
}

func GetUsersCount(w http.ResponseWriter, r *http.Request) {
	collection := db.Instance.Database.Collection("users")

	a, _ := collection.CountDocuments(context.TODO(), bson.D{{}})

	logger.Instance.LogInfo(string(a))

	w.Write([]byte(string(a)))
}
