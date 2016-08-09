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
// START_STRUCT
type output struct {
	Timestamp string `json:"timestamp"`
	TotalRAM  uint64 `json:"totalram"`
	FreeRAM   uint64 `json:"freemem"`
	FreeSWAP  uint64 `json:"freeswap"`
}

// END_STRUCT

func getMem(w http.ResponseWriter, req *http.Request) {
	// Make sure that the writer supports flushing.
	//
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
	enc := json.NewEncoder(w)
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// START_DISPLAY OMIT
		fmt.Fprintf(w, "data: ")

		enc.Encode(&output{time.Now().Format("Mon Jan 02 15:04:05 -0700 2006"), s.Totalram, s.Freeram, s.Freeswap})
		fmt.Fprintf(w, "\n\n")
		// END_DISPLAY OMIT
		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	http.HandleFunc("/mem", getMem) // HL
	http.Handle("/", http.FileServer(http.Dir("htdocs")))
	fmt.Println("serving on http://localhost:7777/mem")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
