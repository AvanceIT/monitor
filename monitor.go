// Monitor provides a client side monitoring suite
package main

import (
	"fmt"
	"monitor/monitors/procmon"
)

func main() {
	if !procmon.RunChecks() {
		fmt.Println("No alerts")
	}
}
