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
	fmt.Println("Total ram:", s.Totalram)
	fmt.Println("Free ram:", s.Freeram)
	// END_DISPLAY OMIT
}
