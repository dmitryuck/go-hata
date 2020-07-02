package routes

import (
	"context"
	"net/http"

	"project/internal/db"
	"project/internal/logger"
	"project/internal/models"
	"project/internal/response"
	"project/internal/services"
	"project/internal/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// ApplyDialogRoutes for methods
func ApplyDialogRoutes(router *mux.Router) {
	//router.HandleFunc(utils.BuildRouteURL(""), GetTest).Methods("GET")
	router.HandleFunc(utils.BuildRouteURL(FetchDialogsRoute), GetFetchDialogs).Methods("GET")
}

// GetTest test route
func GetTest(w http.ResponseWriter, r *http.Request) {
	result := utils.BuildRouteURL(FetchDialogsRoute, "id")
	logger.Instance.LogInfo(result)
}

// GetFetchDialogs router
func GetFetchDialogs(w http.ResponseWriter, r *http.Request) {
	dialog := services.DialogService.FetchDialogs(services.DialogService{})

	dialogResponse := models.DialogResponse.Make(models.DialogResponse{}, &dialog)

	response.ResponseObject.Make(response.ResponseObject{}, w, response.StatusSuccess, dialogResponse)
}

// GetUsersCount users count
func GetUsersCount(w http.ResponseWriter, r *http.Request) {
	collection := db.Instance.Database.Collection("users")

	//a, _ := collection.CountDocuments(context.TODO(), bson.D{{}})
	//logger.Instance.LogInfo(strconv.FormatInt(a, 10))

	var a models.User

	collection.FindOne(context.TODO(), bson.M{"name": "Test"}).Decode(&a)

	w.Write([]byte(string(a.DeviceID)))
}
