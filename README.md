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

## Response

    $ curl -i http://localhost:8080/api/stats/
    HTTP/1.1 200 OK
    Content-Length: 375
    Content-Type: application/json
    Date: Sat, 30 Nov 2013 00:42:54 GMT
    
    {
        "go_version": "1.2.2",
        "go_os": "darwin",
        "go_arch": "amd64",
        "gc_num": 1,
        "gc_last": 1385772060688109000,
        "gc_next": 1622624,
        "memory_lookups": 68,
        "memory_sys": 272289616,
        "memory_total_alloc": 1257976,
        "memory_alloc": 1228864,
        "cgo_call_num": 0,
        "gomaxprocs": 4,
        "goroutine_num": 5,
        "cpu_num": 4,
        "memory_mallocs": 2185,
        "memory_frees": 59,
        "heap_alloc": 1228864,
        "heap_sys": 2097152,
        "heap_idle": 720896,
        "heap_inuse": 1376256,
        "heap_released": 0,
        "heap_objects": 2126
    }

## Plugins

- Zabbix: [fukata/golang-stats-api-handler-zabbix-userparameter](https://github.com/fukata/golang-stats-api-handler-zabbix-userparameter)
- Munin: [fukata/golang-stats-api-handler-munin-plugin](https://github.com/fukata/golang-stats-api-handler-munin-plugin) 
