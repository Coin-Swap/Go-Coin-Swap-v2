package coinswapapi

import (
"net/http"
"io/ioutil"
"fmt"
)

func DialCoinSwapPrivate(url string, cookie *http.Cookie) []byte {
    // Create the client
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    req.AddCookie(cookie)
    Errorcheck(err)
    
    // Make the request
    resp, err := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    Errorcheck(err)
    return body

}

func DialCoinSwapPublic(url string) []byte {
    // Create the client
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    Errorcheck(err)
    
    // Make the request
    resp, err := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    Errorcheck(err)
    return body

}

func Errorcheck(err interface{}){
    if err != nil {
        fmt.Printf("%s",err)
    }
}