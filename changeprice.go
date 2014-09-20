package coinswapapi

import (
"strconv"
"net/http"
"encoding/json"
)

type Change struct{
    orderid string
    marketid string
    direction int
    price string
    amount string
}

// The Sell function will create a market order of type sell. You must pass it the config struct,
// the marketid, the price, the amount, and the cookie.
func ChangePrice(config Config,orderid string,price float64,cookie *http.Cookie) Change {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    url := "https://api.coin-swap.net/order/change/"+orderid+"/"+config.Apikey+"/"+strconv.FormatFloat(price,'f',8,64)
    apiResponse := DialCoinSwapPrivate(url, cookie) 

    // We receive a json response from the server and unmarshal it.
    var change Change
    jsonerr := json.Unmarshal(apiResponse, &change)
    Errorcheck(jsonerr)

    return change
}
