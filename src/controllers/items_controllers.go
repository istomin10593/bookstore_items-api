package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/istomin10593/bookstore_items-api/src/domain/items"
	service "github.com/istomin10593/bookstore_items-api/src/services"
	"github.com/istomin10593/bookstore_items-api/src/utils/http_utils"
	"github.com/istomin10593/bookstore_oauth-go/oauth"
	"github.com/istomin10593/bookstore_utils-go/rest_errors"
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
	if oauthErr := oauth.AuthenticateRequest(r); oauthErr != nil {
		http_utils.ResponseError(w, oauthErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		restErr := rest_errors.NewRestError("invalid request body", http.StatusBadRequest, "invalid request body", []interface{}{err})
		http_utils.ResponseError(w, restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewRestError("invalid item json body", http.StatusBadRequest, "invalid item json body", []interface{}{err})
		http_utils.ResponseError(w, restErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	result, saveErr := service.ItemsService.Create(itemRequest)
	if saveErr != nil {
		http_utils.ResponseError(w, saveErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
