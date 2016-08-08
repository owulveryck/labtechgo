// START_IMPORT OMIT
package main

import (
	"fmt"
	"log"
	"syscall" // HL
)

// START_STRUCT OMIT
func main() {
	//END_IMPORT OMIT
	var s syscall.Sysinfo_t // HLSysinfo
	// START_ERROR OMIT
	err := syscall.Sysinfo(&s) // HLs

	if err != nil {
		log.Fatal(err)
	}
	// END_STRUCT OMIT
	// END_ERROR OMIT
	// START_DISPLAY OMIT
	fmt.Printf("RAM: %v / %v\nSWAP: %v / %v\n", s.Freeram, s.Totalram, s.Freeswap, s.Totalswap)
	// END_DISPLAY OMIT
}
