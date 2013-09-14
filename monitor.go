// Monitor provides a client side monitoring suite
package main

import (
	"fmt"
	"github.com/AvanceIT/monitor/monitors/fsmon"
	"github.com/AvanceIT/monitor/monitors/procmon"
	"github.com/AvanceIT/monitor/monitors/httpmon"
	"github.com/AvanceIT/monitor/tools"
)

func main() {
	tools.Logger("Monitor", "Starting monitor checks")
	if !procmon.RunChecks() {
		fmt.Println("No procmon alerts")
	}

	if !fsmon.RunChecks() {
		fmt.Println("No fsmon alerts")
	}
	// tools.Logger("Monitor", "All checks completed")

	if !httpmon.RunChecks() {
		fmt.Println("No httpmon alerts")
		}
	tools.Logger("Monitor", "All checks completed")
}
