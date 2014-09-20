package main

import (
"strconv"
"math/rand"
"net/http"
)

// This is the format for each open order response
type Openorder struct {
Orderid string
Marketid string
Type string
Balance string
Price string
}
// This will have all of our open orders in a slice
type OrderContainer struct {
Orders []Openorder
}

// This will ask for all open orders
func Openorders(config Config,cookie *http.Cookie) OrderContainer {
    var url string
    url = "https://api.coin-swap.net/order/open/"+config.Apikey
    // Retrieve all open orders.
    response := DialCoinSwap(url, cookie) 
    // Unmarshal the json
    var container = OrderContainer
    err := json.Unmarshal(response, &container)
    Errorcheck(err)
    return container
}
