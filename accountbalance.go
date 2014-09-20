package coinswapapi

import (
"net/http"
"encoding/json"
)

type Balance struct {
Symbol string `json:"symbol"`
Balance string `json:"balance"`
}

// This will retrieve an account balance for the specified currency
func Accountbalance(config Config,cookie *http.Cookie,coin string) Balance {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/balance/"+coin+"/"+config.Apikey
    apiResponse := DialCoinSwapPrivate(url, cookie) 
    
    // We receive a json response from the server and unmarshal it.
    var balance Balance
    jsonerr := json.Unmarshal(apiResponse, &balance)
    Errorcheck(jsonerr)

    return balance
}
