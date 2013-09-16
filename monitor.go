// Monitor provides a client side monitoring suite
package main

import (
	"github.com/AvanceIT/monitor/monitors/fsmon"
	"github.com/AvanceIT/monitor/monitors/httpmon"
	"github.com/AvanceIT/monitor/monitors/logmon"
	"github.com/AvanceIT/monitor/monitors/procmon"
	"github.com/AvanceIT/monitor/tools"
)

func main() {
	tools.Logger("Monitor", "Starting monitor checks")
	if !procmon.RunChecks() {
		tools.Logger("Monitor", "No procmon alerts")
	}

	if !fsmon.RunChecks() {
		tools.Logger("Monitor", "No fsmon alerts")
	}

	if !httpmon.RunChecks() {
		tools.Logger("Monitor", "No httpmon alerts")
	}

	if !logmon.RunChecks() {
		tools.Logger("Monitor", "No logmon alerts")
	}

	tools.Logger("Monitor", "All checks completed")
}
