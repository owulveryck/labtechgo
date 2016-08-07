// +build OMIT
// START_IMPORT OMIT
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"syscall" // HL
	//	"time"
)

// START_STRUCT OMIT
func getMem(w http.ResponseWriter, req *http.Request) {
	//END_IMPORT OMIT
	var s syscall.Sysinfo_t // HLSysinfo
	// START_ERROR OMIT
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// END_STRUCT OMIT
		// END_ERROR OMIT
		// START_DISPLAY OMIT
		enc := json.NewEncoder(w)
		enc.Encode(map[string]uint64{"total": s.Totalram, "free": s.Freeram, "swap": s.Freeswap})
		//		time.Sleep(100 * time.Millisecond)
		// END_DISPLAY OMIT
	}
}

func main() {
	http.HandleFunc("/mem", getMem) // HL
	fmt.Println("serving on http://localhost:7777/mem")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
