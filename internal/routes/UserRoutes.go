package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ApplyUserRoutes export
func ApplyUserRoutes(router *mux.Router) {
	//router.HandleFunc("/", A).GetMethods()
}

// A router
func A(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(string("Hello world")))
}
