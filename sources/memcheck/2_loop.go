package main

// START_IMPORT OMIT
import (
	"fmt"
	"log"
	"os" // To access Stdout // HL
	"syscall"
	"time" // HL
)

// END_IMPORT OMIT

func main() {
	var s syscall.Sysinfo_t
	// START_LOOP OMIT
	for {
		err := syscall.Sysinfo(&s) // HLs

		if err != nil {
			log.Fatal(err)
		}
		// Introducing the Fprintf to write to a "stream" // HL
		w := os.Stdout // HL
		fmt.Fprintf(w, "RAM: %v / %v\nSWAP: %v / %v\n", s.Freeram, s.Totalram, s.Freeswap, s.Totalswap)
		time.Sleep(100 * time.Millisecond) // HL
	}
	// END_LOOP OMIT
}
