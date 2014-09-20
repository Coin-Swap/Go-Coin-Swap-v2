package coinswapapi

import (
"encoding/json"
)

type MarketStats struct {
    Ask string `json:"ask"`
    Bid string `json:"bid"`
    Marketid string `json:"marketid"`
    Symbol string `json:"symbol"`
    Exchange string `json:"exchange"`
    Lastprice string `json:"lastprice"`
    Dayvolume string `json:"dayvolume"`
    Dayhigh string `json:"dayhigh"`
    Daylow string `json:"daylow"`
    Openorders string `json:"openorders"`
}

// The GetMarketStats function will request stats on a specific market.
// The market variable is a string in the form of a pairing.
// Example: var market = "DOGE/BTC"
// You must pass this string variable to the function.
func GetMarketStats(market string) MarketStats {
    // Create our url.
    // Send our URL to the DialCoinSwapPublic function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/market/stats/" + market
    apiResponse := DialCoinSwapPublic(url)

    // We receive a json response from the server and unmarshal it.
    var marketStats MarketStats
    jsonerr := json.Unmarshal(apiResponse, &marketStats)
    Errorcheck(jsonerr)

    return marketStats
}
