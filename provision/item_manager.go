package provision

import (
	"fmt"
	"osrs-items-api/oldschoolrs"
	"osrs-items-api/rsbuddy"
	"osrs-items-api/types"
	"strconv"
	"strings"
)

type ItemManager struct {
	OldschoolRsClient oldschoolrs.OldschoolRsClient
	RSBuddyClient     rsbuddy.RSBuddyClient
}

func (im *ItemManager) FindItemById(itemId string) (itemFull types.RunescapeItem, err error) {
	items, err := im.RSBuddyClient.ListItems()
	if err != nil {
		return itemFull, err
	}

	for _, item := range items {
		if strconv.Itoa(item.ID) == itemId {
			price, err := im.OldschoolRsClient.FindItemPrice(item.ID)
			if err != nil {
				return itemFull, err
			}

			itemFull = types.RunescapeItem{
				Details:        item,
				PriceVariation: price,
			}

			return itemFull, err
		}
	}

	return itemFull, fmt.Errorf("item with id: '%s' was not found", itemId)
}

func (im *ItemManager) FindItemByName(itemName string) (itemsFound []rsbuddy.Item, err error) {
	items, err := im.RSBuddyClient.ListItems()
	if err != nil {
		return itemsFound, err
	}

	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(itemName)) {
			itemsFound = append(itemsFound, item)
		}
	}

	return itemsFound, err
}
