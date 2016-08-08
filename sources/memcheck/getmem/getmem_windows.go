// Package getmem is a cross plateform package to access the memory information via syscall
// +build windows
package getmem

import (
	"fmt"
	"os"
	"syscall" // HL
	"time"
	"unsafe"
)
import (
	"syscall"
	"unsafe"
)

var (
	mod                      = syscall.NewLazyDLL("kernel32.dll")
	procGlobalMemoryStatusEx = mod.NewProc("GlobalMemoryStatusEx")
)

// MEMORYSTATUSEX holds the informations
type MEMORYSTATUSEX struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailageFile         uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

// GetMem returns the memory usage on windows
func GetMem() error {
	var memInfo MEMORYSTATUSEX
	for {
		memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
		mem, _, _ := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
		if mem == 0 {
			return fmt.Errorf("Cannot get information")
		}
		w := os.Stdout
		fmt.Fprintf(w, "RAM: %v / %v\nSWAP: %v / %v\n", memInfo.ullAvailPhys, memInfo.ullTotalPhys)
		time.Sleep(100 * time.Millisecond)
	}
}
