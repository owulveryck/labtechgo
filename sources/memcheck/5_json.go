// +build OMIT
// START_IMPORT OMIT
package main

import (
	"encoding/json" // HL
	"fmt"
	"log"
	"net/http"
	"syscall"
)

// END_IMPORT OMIT

func getMem(w http.ResponseWriter, req *http.Request) {
	var s syscall.Sysinfo_t // HLSysinfo
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// START_DISPLAY OMIT
		enc := json.NewEncoder(w)
		enc.Encode(map[string]uint64{"total": s.Totalram, "free": s.Freeram, "swap": s.Freeswap})
		// END_DISPLAY OMIT
	}
}

func main() {
	http.HandleFunc("/mem", getMem) // HL
	fmt.Println("serving on http://localhost:7777/mem")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
