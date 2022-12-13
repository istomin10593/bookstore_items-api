package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/istomin10593/bookstore_items-api/src/domain/items"
	"github.com/istomin10593/bookstore_items-api/src/domain/queries"
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
	Search(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func getId(r *http.Request) string {
	vars := mux.Vars(r)
	return strings.TrimSpace(vars["id"])
}

func unmarshalRequest(w http.ResponseWriter, r *http.Request, result interface{}) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewRestError("invalid request body", http.StatusBadRequest, "invalid request body", []interface{}{err})
		http_utils.ResponseError(w, restErr)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &result); err != nil {
		restErr := rest_errors.NewRestError("invalid item json body", http.StatusBadRequest, "invalid item json body", []interface{}{err})
		http_utils.ResponseError(w, restErr)
		return
	}
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if oauthErr := oauth.AuthenticateRequest(r); oauthErr != nil {
		http_utils.ResponseError(w, oauthErr)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		restErr := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.ResponseError(w, restErr)
		return
	}

	var itemRequest items.Item
	unmarshalRequest(w, r, &itemRequest)

	itemRequest.Seller = sellerId

	result, saveErr := service.ItemsService.Create(itemRequest)
	if saveErr != nil {
		http_utils.ResponseError(w, saveErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	itemId := getId(r)

	item, err := service.ItemsService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	var query queries.EsQuery
	unmarshalRequest(w, r, &query)

	items, searchErr := service.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.ResponseError(w, searchErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, items)
}

func (c *itemsController) Update(w http.ResponseWriter, r *http.Request) {
	if oauthErr := oauth.AuthenticateRequest(r); oauthErr != nil {
		http_utils.ResponseError(w, oauthErr)
		return
	}

	itemId := getId(r)

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		restErr := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.ResponseError(w, restErr)
		return
	}

	var itemRequest items.Item
	unmarshalRequest(w, r, &itemRequest)

	itemRequest.Seller = sellerId
	itemRequest.Id = itemId

	result, uptErr := service.ItemsService.Update(itemRequest)
	if uptErr != nil {
		http_utils.ResponseError(w, uptErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, result)
}

func (c *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	itemId := getId(r)

	if err := service.ItemsService.Delete(itemId); err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, map[string]string{"status": "deleted"})
}
