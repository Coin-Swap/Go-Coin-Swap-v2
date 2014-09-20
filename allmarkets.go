package coinswapapi

import (
"net/http"
"encoding/json"
)

type Allmarkets struct {
    Id string `json:"id"`
    Symbol string `json:"symbol"`
    Exchange string `json:"exchange"`
    Lastprice string `json:"lastprice"`
    Dayvolume string `json:"dayvolume"`
    Dayhigh string `json:"dayhigh"`
    Daylow string `json:"daylow"`
    Ask string `json:"ask"`
    Bid string `json:"bid"`
    Openorders string `json:"openorders"`
}

// This struct will contain a splice of all our Allmarkets type
type MarketContainer struct {
    Markets []Allmarkets
}

// The Getallmarkets function will request a market summary for every market that Coin-Swap.net has active.
func Getallmarkets(cookie *http.Cookie) MarketContainer {
    // Create our url.
    // Send our URL to the DialCoinSwapPublic function to create an order and receive response.
    var url string
    url = "https://api.coin-swap.net/market/summary/"
    apiResponse := DialCoinSwapPublic(url) 

    // We receive a json response from the server and unmarshal it.
    var marketcontainer MarketContainer
    jsonerr := json.Unmarshal(apiResponse, &marketcontainer)
    Errorcheck(jsonerr)

    return marketcontainer
}
