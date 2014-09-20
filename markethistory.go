package coinswapapi

import (
"encoding/json"
)

type MarketHistory struct {
    Timestamp int32 `json:"timestamp"`
    Price string `json:"price"`
    Amount string `json:"amount"`
}

type MarketHistoryContainer struct {
Orders []MarketHistory
}
// The GetMarketHistory function will request trade history for a specific market.
// The market variable is a string in the form of a pairing.
// Example: var market = "DOGE/BTC"
// You must pass this string variable to the function.
func GetMarketHistory(market string) MarketHistoryContainer {
    // Create our url.
    // Send our URL to the DialCoinSwapPublic function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/market/history/" + market
    apiResponse := DialCoinSwapPublic(url)

    // We receive a json response from the server and unmarshal it.
    var container MarketHistoryContainer
    jsonerr := json.Unmarshal(apiResponse, &container)
    Errorcheck(jsonerr)

    return container
}
