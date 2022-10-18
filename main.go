package main

import (
	"NetworkScan/scan"
	"fmt"
)

func main() {

	fmt.Println("Select protocol:")
	var choice, startPort, endPort int
	fmt.Println("1. TCP")
	fmt.Println("2. UDP")
	fmt.Scanln(&choice)
	fmt.Println("Select start port:")
	fmt.Scanln(&startPort)
	fmt.Println("Select end port:")
	fmt.Scanln(&endPort)
	switch choice {
	case 1:
	}
	scan.PortScanRange(startPort, endPort, "tcp")
}
