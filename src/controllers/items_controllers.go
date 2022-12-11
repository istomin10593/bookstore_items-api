package controllers

import (
	"net/http"

	"github.com/istomin10593/bookstore_items-api/src/domain/items"
	service "github.com/istomin10593/bookstore_items-api/src/services"
	"github.com/istomin10593/bookstore_items-api/src/utils/http_utils"
	"github.com/istomin10593/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := service.ItemsService.Create(item)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
