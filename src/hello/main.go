package main
 
import (
    "log"
    "net/http"
)
 
func main() {
    broken1
    router := NewRouter()
 
    log.Fatal(http.ListenAndServe(":8000", router))
}
