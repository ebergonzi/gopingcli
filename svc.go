package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    response, err := http.Get("http://localhost:8080/")
    if err != nil {
    	log.Fatal(err)
    	return
    } 
    defer response.Body.Close()
    buf, err := ioutil.ReadAll(response.Body)
    bs := string(buf)
    dstUrl := response.Request.URL.String()
    fmt.Fprintf(w, "%s → %s\n", dstUrl, bs)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
