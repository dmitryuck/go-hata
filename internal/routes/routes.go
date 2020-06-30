package routes

import (
	"github.com/gorilla/mux"
)

// ApplyRoutes for root
func ApplyRoutes(router *mux.Router) {
	ApplyDialogRoutes(router)
	ApplyUserRoutes(router)
}
