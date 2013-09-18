/*
Package logmon monitors the provided log files for errors

*/
package logmon

import (
	"bufio"
	"github.com/AvanceIT/monitor/tools"
	"os"
	"strings"
)

var monName string = "logmon"

type configOptions struct {
	Logfiles []string
}

// configMonitor reads the configuration file and sets up
// this monitor. The configuration file just contains
// a list of log files that will be checked 1 per line.
func configMonitor() configOptions {
	var config configOptions
	cl := tools.ReadConfig(monName)
	for _, line := range cl {
		config.Logfiles = append(config.Logfiles, line.Fields[0])
	}
	return config
}

func checkFile(lf string) (alerted bool) {
	tools.Logger(monName, "Checking "+lf)
	lfile, err := os.Open(lf)
	if err != nil {
		tools.Logger(monName, "Unable to open logfile : "+lf)
		return
	}
	defer lfile.Close()

	lfscan := bufio.NewScanner(lfile)
	for lfscan.Scan() {
		tl := lfscan.Text()
		if strings.Contains(tl, "error") {
			alertString := monName + ": Error found in " + lf + " :: " + tl
			tools.RaiseAlert(alertString, 99)
			tools.Logger(monName, "Error found")
			alerted = true
		}
	}

	return
}

func RunChecks() (alerted bool) {
	tools.Logger(monName, "Starting")
	config := configMonitor()

	for _, lf := range config.Logfiles {
		if checkFile(lf) {
			alerted = true
		}
	}

	if alerted {
		tools.Logger(monName, "Alert raised")
	}
	tools.Logger(monName, "Completed")
	return
}
