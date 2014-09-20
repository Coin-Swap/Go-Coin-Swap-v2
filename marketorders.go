package coinswapapi

import (
"encoding/json"
)

type MarketOrder struct {
    Price string `json:"price"`
    Amount string `json:"amount"`
}

type MarketOrderContainer struct {
Orders []MarketOrder
}
// The GetMarketOrders function will retrieve open orders for a specific market.
// The market variable is a string in the form of a pairing.
// Example: var market = "DOGE/BTC"
// You must pass this string variable to the function.
func GetMarketOrders(market string) MarketOrderContainer {
    // Create our url.
    // Send our URL to the DialCoinSwapPublic function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/market/orders/" + market
    apiResponse := DialCoinSwapPublic(url)

    // We receive a json response from the server and unmarshal it.
    var container MarketOrderContainer
    jsonerr := json.Unmarshal(apiResponse, &container)
    Errorcheck(jsonerr)

    return container
}
