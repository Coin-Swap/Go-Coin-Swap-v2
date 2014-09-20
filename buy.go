package coinswapapi

import (
"strconv"
"net/http"
"encoding/json"
)

type Buyorder struct{
    orderid string
    marketid string
    direction int
    price string
    amount string
}

// The Buy function will create a market order of type buy. You must pass it the config struct,
// the marketid, the price, the amount, and the cookie.
func Buy(config Config,marketid string,price float64,amount float64,cookie *http.Cookie) Buyorder {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    url := "https://api.coin-swap.net/order/v2/buy/"+marketid+"/"+config.Apikey+"/"+strconv.FormatFloat(price,'f',8,64)+"/"+strconv.FormatFloat(amount,'f',8,64)
    apiResponse := DialCoinSwapPrivate(url, cookie) 

    // We receive a json response from the server and unmarshal it.
    var buyorder Buyorder
    jsonerr := json.Unmarshal(apiResponse, &buyorder)
    Errorcheck(jsonerr)

    return buyorder
}
