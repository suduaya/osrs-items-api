package types

import (
	"osrs-items-api/oldschoolrs"
	"osrs-items-api/rsbuddy"
)

type RunescapeItem struct {
	Details        rsbuddy.Item                `json:"details"`
	PriceVariation oldschoolrs.ItemPriceRecord `json:"price_history"`
}
