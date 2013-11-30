golang-stats-api-handler
========================

Golang cpu, memory, gc, etc information api handler.

## Install

    go get github.com/fukata/golang-stats-api-handler

## Example

    import (
        "net/http"
        "log"
        "github.com/fukata/golang-stats-api-handler"
    )
    func main() {
        http.HandleFunc("/api/stats", stats_api.Handler)
        log.Fatal( http.ListenAndServe(":8080", nil) )
    }
