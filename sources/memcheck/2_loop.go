// START_IMPORT OMIT
package main

import (
	"fmt"
	"log"
	"os"
	"syscall" // HL
	"time"
)

// START_STRUCT OMIT
func main() {
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
		w := os.Stdout
		fmt.Fprintln(w, "Total ram:", s.Totalram)
		fmt.Fprintln(w, "Free ram:", s.Freeram)
		fmt.Fprintln(w, "Free swap:", s.Freeswap)
		time.Sleep(100 * time.Millisecond)
		// END_DISPLAY OMIT
	}
}
