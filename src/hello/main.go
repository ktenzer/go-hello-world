package main
 
import (
    "log"
    "net/http"
)
 
func main() {
    broken
    router := NewRouter()
 
    log.Fatal(http.ListenAndServe(":8000", router))
}
