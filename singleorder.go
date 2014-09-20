package coinswapapi

import (
"net/http"
"encoding/json"
)

type Singleorder struct {
Orderid string `json:"orderid"`
Marketid string `json:"marketid"`
Type string `json:"type"`
Balance string `json:"balance"`
Price string `json:"price"`
}

// This will retrieve information for the specified order
func GetSingleorder(config Config,cookie *http.Cookie,orderid string) Singleorder {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/order/single/"+orderid+"/"+config.Apikey
    apiResponse := DialCoinSwapPrivate(url, cookie) 
    
    // We receive a json response from the server and unmarshal it.
    var singleorder Singleorder
    jsonerr := json.Unmarshal(apiResponse, &singleorder)
    Errorcheck(jsonerr)

    return singleorder
}
