package main

import (
"strconv"
"math/rand"
"net/http"
)

type Canceled struct {
Orderid string
Balance string
Response string
}
// This will create a sell order
func Cancelorder(config Config,cookie *http.Cookie,openorder Openorder) Order {
    var url string
    url = "https://api.coin-swap.net/order/cancel/"+openorder.Orderid+"/"+config.Apikey
    response := DialCoinSwap(url, cookie) // Send off our URL to create an order. Receive response.
    
    var canceled Canceled
    jsonerr := json.Unmarshal(response, &canceled)
    Errorcheck(jsonerr)

    return canceled
}
