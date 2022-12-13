package items

import (
	"github.com/istomin10593/bookstore_utils-go/rest_errors"
)

type Item struct {
	Id                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Descripton        Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlaintText string `json:"plain_text"`
	Html       string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

func (i *Item) Validate() rest_errors.RestErr {
	if i.Title == "" {
		return rest_errors.NewBadRequestError("invalid book title")
	}
	if i.Descripton.PlaintText == "" {
		return rest_errors.NewBadRequestError("invalid book plaint text")
	}
	if i.Price <= 0 {
		return rest_errors.NewBadRequestError("invalid book price")
	}
	if i.Status == "" {
		return rest_errors.NewBadRequestError("invalid book status")
	}
	return nil
}
