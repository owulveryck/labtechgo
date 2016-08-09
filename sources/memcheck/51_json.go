// +build OMIT
// START_IMPORT OMIT
package main

import (
	"encoding/json" // HL
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"
)

// END_IMPORT OMIT

// START_STRUCT OMIT
type result struct {
	Timestamp string `json:"timestamp"`
	TotalRAM  uint64 `json:"totalram"`
	FreeRAM   uint64 `json:"freeram"`
	Swap      uint64 `json:"freeswap"`
}

// END_STRUCT OMIT

func getMem(w http.ResponseWriter, req *http.Request) {
	// Make sure that the writer supports flushing.
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var s syscall.Sysinfo_t // HLSysinfo
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// START_DISPLAY OMIT
		fmt.Fprintf(w, "data: ")
		var r result
		r.Timestamp = time.Now().Format("Mon Jan 02 15:04:05 -0700 2006")
		r.TotalRAM = s.Totalram
		r.FreeRAM = s.Freeram
		r.Swap = s.Freeswap
		b, err := json.Marshal(r) // HL
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(b))
		// END_DISPLAY OMIT
		flusher.Flush()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	http.HandleFunc("/mem", getMem) // HL
	fmt.Println("serving on http://localhost:7777/mem")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
