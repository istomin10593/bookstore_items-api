package items

import (
	"errors"

	"github.com/istomin10593/bookstore_items-api/src/clients/elasticsearch"
	"github.com/istomin10593/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("errors when trying to save item", errors.New("database error"))
	}

	i.Id = result.Id

	return nil
}
