package scan

import (
	"fmt"
	"net"
	"strconv"
	_ "time"
)

func PortScan(port int, protocol string) {

	server := "localhost"
	address := server + ":" + strconv.Itoa(port)
	_, err := net.Dial(protocol, address)
	if err == nil {
		fmt.Printf("Port:%d\t%s\t%s\t\n", port, "open", protocol)
	} else {
		fmt.Printf("Port:%d\t%s\t%s\t\n", port, "closed", protocol)
	}
}
func PortScanRange(startPort int, endPort int, protocol string) {
	for i := startPort; i <= endPort; i++ {
		PortScan(i, protocol)
	}
}
