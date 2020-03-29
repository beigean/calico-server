// ref. https://go-tour-jp.appspot.com/methods/18
package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (addr *IPAddr) String() string {
	// return string(addr[0]) + "." + string(addr[1]) + "." + string(addr[2]) + "." + string(addr[3])
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

func main() {
	var i fmt.Stringer
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		i = &ip
		fmt.Printf("%v: %v\n", name, i.String())
	}
}
