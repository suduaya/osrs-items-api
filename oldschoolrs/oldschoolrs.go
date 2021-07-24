package oldschoolrs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Oldschool Runescape Client
type OldschoolRsClient struct {
	*http.Client
	Host string
}

func New(apiUrl string) OldschoolRsClient {
	return OldschoolRsClient{
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

type ItemPriceRecord struct {
	Daily   []Price `json:"daily"`
	Average []Price `json:"average"`
}

type Price struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
}

func (osrs *OldschoolRsClient) FindItemPrice(itemId int) (record ItemPriceRecord, err error) {
	resp, err := osrs.Get(osrs.Host + fmt.Sprintf("/%d.json", itemId))
	if err != nil {
		return record, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return record, fmt.Errorf("failed to list item price record, got response status: '%d'", resp.StatusCode)
	}

	var itemPricesToMap map[string]interface{}          // Response format is not a json array
	json.NewDecoder(resp.Body).Decode(&itemPricesToMap) // We need to map all the items first

	daily := itemPricesToMap["daily"].(map[string]interface{}) // Save daily price records
	for dailyTimestamp, dailyValue := range daily {
		record.Daily = append(record.Daily, Price{
			Timestamp: dailyTimestamp,
			Value:     dailyValue.(float64),
		})
	}

	average := itemPricesToMap["average"].(map[string]interface{}) // Save average price records
	for averageTimestamp, averageValue := range average {
		record.Average = append(record.Average, Price{
			Timestamp: averageTimestamp,
			Value:     averageValue.(float64),
		})
	}

	return record, err
}
