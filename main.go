package main

import (
	"os"
	"osrs-items-api/api"
	"osrs-items-api/oldschoolrs"
	"osrs-items-api/rsbuddy"
)

func main() {
	port := os.Getenv("PORT")

	rsbuddyClient := rsbuddy.New("https://rsbuddy.com/exchange/summary.json")
	osClient := oldschoolrs.New("https://secure.runescape.com/m=itemdb_oldschool/api/graph")

	itemsApi := api.New(osClient, rsbuddyClient, port)
	itemsApi.Start()
}
