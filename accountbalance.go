package coinswapapi

import (
"strconv"
"math/rand"
"net/http"
)

type Balance struct {
symbol string
balance string
}

// This will retrieve an account balance for the specified currency
func Accountbalance(config Config,cookie *http.Cookie,coin string) Balance {
    var url string
    url = "https://api.coin-swap.net/balance/"+coin+"/"+config.Apikey
    response := DialCoinSwap(url, cookie) // Send off our URL to create an order. Receive response.
    
    var balance Balance
    jsonerr := json.Unmarshal(response, &balance)
    Errorcheck(jsonerr)

    return balance
}
