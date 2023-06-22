package service

import (
	"fmt"
	"osrs-items-api/pkg/oldschoolrs"
)

type Service struct {
	OldschoolRsClient *oldschoolrs.OldschoolRsClient
}

func New(osrs *oldschoolrs.OldschoolRsClient) *Service {
	return &Service{
		OldschoolRsClient: osrs,
	}
}

func (s *Service) FindItemById(itemId int) (itemFull RunescapeItem, err error) {

	price, err := s.OldschoolRsClient.FindItemPrice(itemId)
	if err != nil {
		return itemFull, fmt.Errorf("item with id: '%d' was not found", itemId)
	}

	itemFull = RunescapeItem{
		PriceVariation: price,
	}

	return itemFull, nil
}
