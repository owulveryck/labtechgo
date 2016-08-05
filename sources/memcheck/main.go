package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	var s syscall.Sysinfo_t
	err := syscall.Sysinfo(&s)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total ram:", s.Totalram)
	fmt.Println("Free ram:", s.Freeram)
}
