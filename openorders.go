package coinswapapi

import (

"encoding/json"
"net/http"
)

// This is the format for each open order response
type Openorder struct {
Orderid string `json:"orderid"`
Marketid string `json:"marketid"`
Type string `json:"type"`
Balance string `json:"balance"`
Price string `json:"price"`
}
// This will have all of our open orders in a slice
type OrderContainer struct {
Orders []Openorder
}

// The Openorders function will request all open orders for this user.
func GetOpenorders(config Config,cookie *http.Cookie) OrderContainer {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/order/open/"+config.Apikey
    apiResponse := DialCoinSwapPrivate(url, cookie) 

    // We receive a json response from the server and unmarshal it.
    //var container OrderContainer
    var container []Openorder
    err := json.Unmarshal(apiResponse, &container)
    Errorcheck(err)
    return container
}
