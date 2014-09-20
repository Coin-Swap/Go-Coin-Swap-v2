package coinswapapi

import (
"strconv"
"math/rand"
"net/http"
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

// This will create a buy order
func GetMarketStats(config Config,market string,price float64,amount float64,cookie *http.Cookie) MarketStats {
    //https://api.coin-swap.net/order/v2/buy/{marketid}/{apikey}/{price}/{amount}
    var url string
    url := "https://api.coin-swap.net/market/stats/" + market
    apiResponse := DialCoinSwap(url, cookie) // Send off our URL to create an order. Receive response.

    var marketStats MarketStats
    jsonerr := json.Unmarshal(response, &marketStats)
    Errorcheck(jsonerr)

    return marketStats
}
