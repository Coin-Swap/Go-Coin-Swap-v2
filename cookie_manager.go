package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "crypto/tls"
    "os"
    "bufio"
    "strings"
)


// Borrowed from: http://play.golang.org/p/YkW_z2CSyE
// Turns a raw cookie string into an http.Cookie struct
func Parse_cookie(cookie string) (*http.Cookie, error) {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", cookie))))
	if err != nil {
		return nil, err
	}
	cookies := req.Cookies()
	if len(cookies) == 0 {
		return nil, fmt.Errorf("no cookies")
	}
	return cookies[0], nil
}

// Save a cookie to disk.
func Save_cookie(c *http.Cookie) {

    c_filename := "userkey.cookie"
    fmt.Printf("Saving cookie to disk (%v): %v\n", c_filename, c.Raw)
    ioutil.WriteFile(c_filename, []byte(c.Raw), 0600)
}

// Get a new userkey cookie and save it to disk.
func Get_userkey_cookie(activation_code string) *http.Cookie {

    var result *http.Cookie

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    url := "https://api.coin-swap.net/initialize/" + activation_code

    // Request the url response which will be in the form of json
    urlResponse, urlError := client.Get(url)

    // Check if there was an error:
    if urlError != nil {
        fmt.Printf("Url Error: %v\n", urlError)
        return result
    }
    
    apiResponse,apiError := ioutil.ReadAll(urlResponse.Body)
    urlResponse.Body.Close() // Close the url request
    // Check if there was an error
    if apiError != nil {
        fmt.Printf("API Error: %v\n", apiError)
        return result
    }
    _ = apiResponse
    //fmt.Printf("API Response: %v\n", apiResponse)
    
    fmt.Printf("Cookies:\n")
    
    got_userkey := false
    var userkey *http.Cookie
    for c_num, c := range urlResponse.Cookies() {
        fmt.Printf("%v:\n  Name: %v\n  Value: %v\n\n",c_num, c.Name, c.Value)
        
        if c.Name == "userkey" {
            got_userkey = true
            userkey = c
            break
        }
    }
    
    if !got_userkey {
        fmt.Printf("Cookie Error: Didn't get a userkey cookie from server.\n")
        return result
    }
    
    result = userkey
    Save_cookie(userkey)
    
    return result
}

// Load a cookie from disk.
func Load_cookie(c_filename string) *http.Cookie {
    var result *http.Cookie
    
    raw_cookie, file_error := ioutil.ReadFile(c_filename)
    if file_error != nil {
        fmt.Printf("File Error: %v\n", file_error)
        return result
    }
    
    parsed_cookie, parse_error := Parse_cookie(string(raw_cookie))
    if parse_error != nil {
        fmt.Printf("Cookie Parse Error: %v\n", parse_error)
        return result
    }
    
    result = parsed_cookie    
    fmt.Printf("userkey cookie loaded\n")
    
    return result
}

// Get a new userkey cookie or load one from disk if available. Returns the cookie.
func Load_userkey_cookie(config Config) *http.Cookie {

    var result *http.Cookie

    // If we don't have a userkey cookie on disk.
    if _, err := os.Stat("userkey.cookie"); os.IsNotExist(err) {
    
        fmt.Printf("Userkey cookie not found. Attempting to get a new one.")
        
        if config.Initkey == "" {
            fmt.Printf("Config Error: You need to specify your activation key in the config file's initkey section.\n")
            return result
        }
        
        result = Get_userkey_cookie(config.Initkey) // Get a new userkey cookie and save it to disk.
        
    } else {
        
        result = Load_cookie("userkey.cookie")  // Load a new cookie with the data that was saved in userkey.cookie
    }
    
    return result
}