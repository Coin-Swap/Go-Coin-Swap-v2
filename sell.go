package coinswapapi

import (
"strconv"
"encoding/json"
"net/http"
)

type Sellorder struct{
    Orderid string
    Marketid string
    Direction int
    Price string
    Amount string
}

// The Sell function will create a market order of type sell. You must pass it the config struct,
// the marketid, the price, the amount, and the cookie.
func Sell(config Config,marketid string,price float64,amount float64,cookie *http.Cookie) Sellorder {
    // Create our url.
    // Send our URL to the DialCoinSwapPrivate function to create an order and receive response.
    url := "https://api.coin-swap.net/order/v2/sell/"+marketid+"/"+config.Apikey+"/"+strconv.FormatFloat(price,'f',8,64)+"/"+strconv.FormatFloat(amount,'f',8,64)
    apiResponse := DialCoinSwapPrivate(url, cookie) 

    // We receive a json response from the server and unmarshal it.
    var sellorder Sellorder
    jsonerr := json.Unmarshal(apiResponse, &sellorder)
    Errorcheck(jsonerr)

    return sellorder
}
