package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10*iota) 
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 1024.
	fmt.Printf("%.2fGB\n", fileSize/KB)
}