// Monitor provides a client side monitoring suite
package main

//BUG(JP) Needs an exception handler if the config files are not present :) 

import (
	"fmt"
	"github.com/AvanceIT/monitor/monitors/procmon"
	"github.com/AvanceIT/monitor/monitors/fsmon"
)

func main() {
	if !procmon.RunChecks() {
		fmt.Println("No procmon alerts")
	}

	if !fsmon.RunChecks() {
		fmt.Println("No fsmon alerts")
	}
}
