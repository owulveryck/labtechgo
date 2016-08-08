// Package getmem is a cross plateform package to access the memory information via syscall
// +build linux
package getmem

import (
	"fmt"
	"os"
	"syscall" // HL
	"time"
)

// GetMem displays the memory on a linux system
func GetMem() error { // HL
	var s syscall.Sysinfo_t
	for {
		err := syscall.Sysinfo(&s)

		if err != nil {
			return err
		}
		w := os.Stdout
		fmt.Fprintf(w, "RAM: %v / %v\nSWAP: %v / %v\n", s.Freeram, s.Totalram, s.Freeswap, s.Totalswap)
		time.Sleep(100 * time.Millisecond)
	}
}
