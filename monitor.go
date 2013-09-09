// Monitor provides a client side monitoring suite
package main

import (
	"fmt"
	"monitor/monitors/procmon"
	"monitor/monitors/discmon"
)

func main() {
	if !procmon.RunChecks() {
		fmt.Println("No procmon alerts")
	}

	if !discmon.RunChecks() {
		fmt.Println("No discmon alerts")
	}
}
