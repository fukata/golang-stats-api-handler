package stats_api

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime"
	"strconv"
)

type result struct {
	CpuNum       int   `json:"cpu_num"`
	GoroutineNum int   `json:"goroutine_num"`
	Gomaxprocs   int   `json:"gomaxprocs"`
	CgoCallNum   int64 `json:"cgo_call_num"`
	// memory
	MemoryAlloc      uint64 `json:"memory_alloc"`
	MemoryTotalAlloc uint64 `json:"memory_total_alloc"`
	MemorySys        uint64 `json:"memory_sys"`
	MemoryLookups    uint64 `json:"memory_lookups"`
	MemoryMallocs    uint64 `json:"memory_mallocs"`
	MemoryFrees      uint64 `json:"memory_frees"`
	// heap
	HeapAlloc    uint64 `json:"heap_alloc"`
	HeapSys      uint64 `json:"heap_sys"`
	HeapIdle     uint64 `json:"heap_idle"`
	HeapInuse    uint64 `json:"heap_inuse"`
	HeapReleased uint64 `json:"heap_released"`
	HeapObjects  uint64 `json:"heap_objects"`
	// gabarage collection
	GcNext uint64 `json:"gc_next"`
	GcLast uint64 `json:"gc_last"`
	GcNum  uint32 `json:"gc_num"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	result := &result{
		CpuNum:       runtime.NumCPU(),
		GoroutineNum: runtime.NumGoroutine(),
		Gomaxprocs:   runtime.GOMAXPROCS(0),
		CgoCallNum:   runtime.NumCgoCall(),
		// memory
		MemoryAlloc:      mem.Alloc,
		MemoryTotalAlloc: mem.TotalAlloc,
		MemorySys:        mem.Sys,
		MemoryLookups:    mem.Lookups,
		MemoryMallocs:    mem.Mallocs,
		MemoryFrees:      mem.Frees,
		// heap
		HeapAlloc:    mem.HeapAlloc,
		HeapSys:      mem.HeapSys,
		HeapIdle:     mem.HeapIdle,
		HeapInuse:    mem.HeapInuse,
		HeapReleased: mem.HeapReleased,
		HeapObjects:  mem.HeapObjects,
		// gabarage collection
		GcNext: mem.NextGC,
		GcLast: mem.LastGC,
		GcNum:  mem.NumGC,
	}

	jsonBytes, jsonErr := json.Marshal(result)
	var body string
	if jsonErr != nil {
		body = jsonErr.Error()
	} else {
		body = string(jsonBytes)
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "text/html; charset=UTF-8"
	headers["Content-Length"] = strconv.Itoa(len(body))
	for name, value := range headers {
		w.Header().Set(name, value)
	}

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	io.WriteString(w, body)
}
