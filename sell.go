package main

import (
"strconv"
"math/rand"
"net/http"
)

type Sellorder struct{
    orderid string
    marketid string
    direction int
    price string
    amount string
}

// The Sell function will create a market order of type sell. You must pass it the config struct,
// the marketid, the price, the amount, and the cookie.
func Sell(config Config,marketid string,price float64,amount float64,cookie *http.Cookie) Sellorder {
    url := "https://api.coin-swap.net/order/v2/sell/"+config.Marketid+"/"+config.Apikey+"/"+strconv.FormatFloat(price,'f',8,64)+"/"+strconv.FormatFloat(amount,'f',8,64)
    // Send our URL to the DialCoinSwap function to create an order and receive response.
    apiResponse := DialCoinSwap(url, cookie) 

    // We receive a json response from the server. We will unmarshal it into our Sellorder struct.
    var sellorder Sellorder
    jsonerr := json.Unmarshal(response, &sellorder)
    Errorcheck(jsonerr)

    return sellorder
}
