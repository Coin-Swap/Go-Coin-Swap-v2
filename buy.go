package coinswapapi

import (
"strconv"
"math/rand"
"net/http"
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
func Buy(config Config,price float64,amount float64,cookie *http.Cookie) Buyorder {
    url := "https://api.coin-swap.net/order/v2/buy/"+config.Marketid+"/"+config.Apikey+"/"+strconv.FormatFloat(price,'f',8,64)+"/"+strconv.FormatFloat(amount,'f',8,64)
    // Send our URL to the DialCoinSwap function to create an order and receive response.
    apiResponse := DialCoinSwap(url, cookie) 

    // We receive a json response from the server. We will unmarshal it into our Buyorder struct.
    var buyorder Buyorder
    jsonerr := json.Unmarshal(response, &buyorder)
    Errorcheck(jsonerr)

    return buyorder
}
