package service

import (
	"github.com/istomin10593/bookstore_items-api/src/domain/items"
	"github.com/istomin10593/bookstore_items-api/src/domain/queries"
	"github.com/istomin10593/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
	Update(items.Item) (*items.Item, rest_errors.RestErr)
	Delete(string) rest_errors.RestErr
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Validate(); err != nil {
		return nil, err
	}

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemsService) Update(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Validate(); err != nil {
		return nil, err
	}

	if err := item.Update(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Delete(id string) rest_errors.RestErr {
	item := items.Item{Id: id}
	return item.Delete()
}
