package coinswapapi

import (
"net/http"
"encoding/json"
)

type Canceled struct {
Orderid string
Balance string
Response string
}
// This will create a sell order
func Cancelorder(config Config,cookie *http.Cookie,openorder Openorder) Canceled {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/order/cancel/"+openorder.Orderid+"/"+config.Apikey
    apiResponse := DialCoinSwapPrivate(url, cookie) 
    
    // We receive a json response from the server and unmarshal it.
    var canceled Canceled
    jsonerr := json.Unmarshal(apiResponse, &canceled)
    Errorcheck(jsonerr)

    return canceled
}
