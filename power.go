package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    
    username := "" 
    password := ""

    
    data := url.Values{}
    data.Set("username", username)
    data.Set("password", password)
    
    req, err := http.NewRequest("POST", "https://www.instagram.com/accounts/login/ajax/", bytes.NewBufferString(data.Encode()))
    if err != nil {
        fmt.Println("Err:", err)
        return
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("User-Agent", "Mozilla/5.0")
    req.Header.Set("X-CSRFToken", "") // CSRF

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Err:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Err", err)
        return
    }

    
    if resp.StatusCode == 200 {
        fmt.Println("Done")
    } else {
        fmt.Println("get out")
    }
}
