package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
)



type Order struct {
    Order string `json:"orderid"`
    Marketid string `json:"marketid"`
    Direction int `json:"direction"`
    Price string `json:"price"`
    Amount string `json:"amount"`
}

// returns the golang equivalent of { "market_name": ask_price, ... }
func Get_market_data(config Config) MarketStats {
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    url := "https://api.coin-swap.net/market/stats/" + config.Market
    

    // Request the url response which will be in the form of json
    urlResponse, urlError := client.Get(url)

    // Check if there was an error:
    if urlError != nil {
            fmt.Printf("%s",urlError)
    }

    // Load the reponse
    apiResponse,apiError := ioutil.ReadAll(urlResponse.Body)
    urlResponse.Body.Close() // Close the url request
    // Check if there was an error
    if apiError != nil {
        fmt.Printf("%s",apiError)
    }
    
    var marketstats MarketStats

    err := json.Unmarshal(apiResponse, &marketstats)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    fmt.Printf("marketstats: %v\n", marketstats)

    return marketstats
}
