package service

import (
	"osrs-items-api/pkg/oldschoolrs"
)

type RunescapeItem struct {
	PriceVariation oldschoolrs.ItemPriceRecord `json:"price_history"`
}
