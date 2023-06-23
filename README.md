# osrs-items-api

Oldschool Runescape Items API

## Description

<p align="center">
<img width="500" height="300" src="https://github.com/suduaya/osrs-items-api/blob/master/docs/images/ge.png?raw=true"/>
</p>

This API extracts data from OSRS Grand Exchange and provides all the pricing information regarding a Runescape Item.

## Supported Operations
  ```
  GET /v1/items/{id}/price
  ```
  ```json
  {
    "price_history": {
      "daily": [
        {
          "timestamp": "1615593600000",
          "value": 302
        }
      ],
      "average": [
        {
          "timestamp": "1615593600000",
          "value": 302
        }
      ]
    }
  }
  ```
