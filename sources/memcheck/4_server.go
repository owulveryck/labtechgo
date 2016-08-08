// +build OMIT
package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall" // HL
)

// START_SIGNATURE OMIT
func getMem(w http.ResponseWriter, req *http.Request) {
	//END_SIGNATURE OMIT
	var s syscall.Sysinfo_t // HLSysinfo
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// START_WRITE OMIT
		fmt.Fprintf(w, "RAM: %v / %v\nSWAP: %v / %v\n", s.Freeram, s.Totalram, s.Freeswap, s.Totalswap) // HLw
		// END_WRITE OMIT
	}
}

// START_HTTP OMIT
func main() {
	http.HandleFunc("/mem", getMem) // HL
	fmt.Println("serving on http://localhost:7777/mem")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

// END_HTTP OMIT
