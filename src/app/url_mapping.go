package app

import (
	"net/http"

	"github.com/istomin10593/bookstore_items-api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
