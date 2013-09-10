// Monitor provides a client side monitoring suite
package main

import (
	"fmt"
	"monitor/monitors/procmon"
	"monitor/monitors/fsmon"
)

func main() {
	if !procmon.RunChecks() {
		fmt.Println("No procmon alerts")
	}

	if !fsmon.RunChecks() {
		fmt.Println("No fsmon alerts")
	}
}
