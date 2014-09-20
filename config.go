package coinswapapi

import (
"os"
"fmt"
"io/ioutil"
"encoding/json"
)

type Config struct {
    Apikey string `json:"apikey"`
    Initkey string `json:"initkey"`
}

func Init_config() Config {
    configfile, err := ioutil.ReadFile("go_bot.cfg")
    if err != nil {
        create_config()
    }
    var config Config
    jsonerr := json.Unmarshal(configfile, &config)
    if jsonerr != nil {
        fmt.Printf("Error: %v\n", jsonerr)
    }
    return config
}

func create_config() {
    // Create our config template and encode it to json
    var config Config
    jsoncfg,err := json.Marshal(config)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }

    // Create the config file
    file, err := os.Create("go_bot.cfg")
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()

    // Write the config to the file
    n2,err :=  file.Write(jsoncfg)
    fmt.Printf("wrote %d bytes\n", n2)

}