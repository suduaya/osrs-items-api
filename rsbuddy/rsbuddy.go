package rsbuddy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

type RSBuddyClient struct {
	*http.Client
	Host string
}

func New(apiUrl string) RSBuddyClient {
	return RSBuddyClient{
		Host: apiUrl,
		Client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
	}
}

type Item struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Members         bool    `json:"members"`
	Sp              float32 `json:"sp"`
	BuyAverage      float32 `json:"buy_average"`
	BuyQuantity     float32 `json:"buy_quantity"`
	SellAverage     float32 `json:"sell_average"`
	SellQuantity    float32 `json:"sell_quantity"`
	OverallAverage  float32 `json:"overall_average"`
	OverallQuantity float32 `json:"overall_quantity"`
}

func (rs *RSBuddyClient) ListItems() (items []Item, err error) {
	resp, err := rs.Get(rs.Host)
	if err != nil {
		return items, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return items, fmt.Errorf("failed to list items, got response status: '%d'", resp.StatusCode)
	}

	var itemsToMap map[string]Item                 // Response format is not a json array
	json.NewDecoder(resp.Body).Decode(&itemsToMap) // We need to map all the items first

	for _, itemMap := range itemsToMap {
		var item Item
		if err := mapstructure.Decode(itemMap, &item); err != nil { // Decode each one to Items struct
			return []Item{}, err
		}
		items = append(items, item) // Append to items
	}

	return items, err
}
