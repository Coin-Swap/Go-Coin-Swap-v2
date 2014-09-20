package main

import (
    "fmt"
)


func main() {
    // Check for a config file. If there is no
    // config file, create a blank template.
    config := Init_config()
    // The following will check for the required cookie.  
    // If no cookie is found it will attempt toretrieve 
    // it from the server. 
    // The config file must be filled out first.
    cookie := Load_userkey_cookie(config)

    _ = Tradebot(config,cookie)

}



