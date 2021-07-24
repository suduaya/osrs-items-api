# osrs-items-api

Oldschool Runescape Items API
## Description

<p align="center">
<img width="500" height="300" src="https://github.com/suduaya/osrs-items-api/blob/master/images/ge.png?raw=true"/>
</p>

This API extracts data from OSRS Grand Exchange and OSBuddy Exchange Tool and provides all the pricing information regarding a Runescape Item.

## Supported Operations 
* Find item by item-id
    * Item Details (with the latest price information)
    * Price Variation
    ```
    GET /items/2
    ```
    ```json
   {
      "details":{
         "id":2,
         "name":"Cannonball",
         "members":true,
         "sp":5,
         "buy_average":210,
         "buy_quantity":3694,
         "sell_average":209,
         "sell_quantity":37208,
         "overall_average":209,
         "overall_quantity":40902
      },
      "price_history":{
         "daily":[
          {
            "timestamp":"1615593600000",
            "value":302
          },
         ...
         ],
         "average":[
         {
           "timestamp":"1615593600000",
           "value":302
         },
         ...
         ]
      }
   }
   ```
* List items (use *name* param for filtering)
 ```
 GET /items?name=yew lo
 ```
 ```json
 [
   {
      "id":855,
      "name":"Yew longbow",
      "members":true,
      "sp":1280,
      "buy_average":0,
      "buy_quantity":0,
      "sell_average":0,
      "sell_quantity":0,
      "overall_average":0,
      "overall_quantity":0
   },
   {
      "id":66,
      "name":"Yew longbow (u)",
      "members":true,
      "sp":640,
      "buy_average":0,
      "buy_quantity":0,
      "sell_average":0,
      "sell_quantity":0,
      "overall_average":0,
      "overall_quantity":0
   },
   {
      "id":1515,
      "name":"Yew logs",
      "members":false,
      "sp":160,
      "buy_average":0,
      "buy_quantity":0,
      "sell_average":210,
      "sell_quantity":80,
      "overall_average":210,
      "overall_quantity":80
   }
]
 ```

