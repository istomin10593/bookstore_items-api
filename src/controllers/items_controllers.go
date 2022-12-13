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
}

type itemsController struct{}

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

	requestBody, err := io.ReadAll(r.Body)
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

	itemRequest.Seller = sellerId

	result, saveErr := service.ItemsService.Create(itemRequest)
	if saveErr != nil {
		http_utils.ResponseError(w, saveErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := service.ItemsService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json")
		http_utils.ResponseError(w, apiErr)
		return
	}

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

	vars := mux.Vars(r)

	itemId := strings.TrimSpace(vars["id"])

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		restErr := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.ResponseError(w, restErr)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
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

	itemRequest.Seller = sellerId
	itemRequest.Id = itemId

	result, uptErr := service.ItemsService.Update(itemRequest)
	if uptErr != nil {
		http_utils.ResponseError(w, uptErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, result)
}
