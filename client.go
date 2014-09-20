package coinswapapi

import (
"net/http"
"io/ioutil"
"fmt"
"encoding/json"
)

func DialCoinSwap(url string, cookie *http.Cookie) []byte {
    
    client := &http.Client{}
    fmt.Printf("url: %v\n", url)
    //resp, err := client.Get(url)
    req, err := http.NewRequest("GET", url, nil)
    req.AddCookie(cookie)
    Errorcheck(err)
    //req.Header.Add("If-None-Match", `W/"wyzzy"`)
    
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