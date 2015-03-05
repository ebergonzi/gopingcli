package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    dst := params["dst"]
    log.Println(dst)
    response, err := http.Get("http://" + dst)
    
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
    rtr := mux.NewRouter()
    rtr.HandleFunc("/v1/ping/{dst:.*}", handler).Methods("GET")

    http.Handle("/", rtr)
    log.Println("up")
    http.ListenAndServe(":8081", nil)
}
