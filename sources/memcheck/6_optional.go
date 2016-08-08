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
	Timestamp time.Time `json:"timestamp"`
	FreeRAM   uint64    `json:"freemem"`
	FreeSWAP  uint64    `json:"freeswap"`
}

// END_STRUCT

func getMem(w http.ResponseWriter, req *http.Request) {
	var s syscall.Sysinfo_t // HLSysinfo
	fmt.Fprintf(w, "[")
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// START_DISPLAY OMIT
		enc := json.NewEncoder(w)

		enc.Encode(&output{time.Now(), s.Freeram, s.Freeswap})
		fmt.Fprintf(w, ",")
		// END_DISPLAY OMIT
	}
	fmt.Fprintf(w, "]")
}

func main() {
	http.HandleFunc("/mem", getMem) // HL
	http.Handle("/", http.FileServer(http.Dir("htdocs")))
	fmt.Println("serving on http://localhost:7777/mem")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
