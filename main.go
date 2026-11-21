package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Hallo nama saya Rida dari jakarta")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9004")
http.ListenAndServe(":9004", nil)

}
