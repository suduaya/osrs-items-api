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

func (s *Service) GetItemPriceVariationById(itemId int) (oldschoolrs.ItemPriceRecord, error) {

	price, err := s.OldschoolRsClient.FindItemPrice(itemId)
	if err != nil {
		return price, fmt.Errorf("item with id: '%d' was not found", itemId)
	}

	return price, nil
}
