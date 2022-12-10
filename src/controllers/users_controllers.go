package controllers

import (
	"fmt"
	"net/http"

	"github.com/istomin10593/bookstore_items-api/src/domain/items"
	service "github.com/istomin10593/bookstore_items-api/src/services"
	"github.com/istomin10593/bookstore_oauth-go/oauth"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type usersController struct{}

func (c *usersController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := service.ItemsService.Create(item)
	if err != nil {
		return
	}

	fmt.Println(result)
}

func (c *usersController) Get(w http.ResponseWriter, r *http.Request) {

}
