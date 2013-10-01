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
var alerted bool = false

type configOptions struct {
	Logfiles []string
}

type logFile struct {
	Filename string
	Errors   bool
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

func (lf logFile) checkFile() logFile {
	lfile, err := os.Open(lf.Filename)
	if err != nil {
		tools.Logger(monName, "Unable to open logfile : "+lf.Filename)
		return lf
	}
	defer lfile.Close()

	lfscan := bufio.NewScanner(lfile)
	for lfscan.Scan() {
		tl := lfscan.Text()
		if strings.Contains(tl, "error") {
			alertString := monName + ": Error found in " +
				lf.Filename + " :: " + tl
			tools.RaiseAlert(alertString, 99)
			tools.Logger(monName, "Error found")
			lf.Errors = true
		}
	}
	return lf
}

func checker(queue <-chan logFile, done chan<- int, alert chan<- int) {
	lfile := <-queue
	message := "Checking " + lfile.Filename
	tools.Logger(monName, message)
	lfile = lfile.checkFile()
	if lfile.Errors {
		alert <- 1
	}
	done <- 1
}

func RunChecks() (alerted bool) {
	tools.Logger(monName, "Starting")
	config := configMonitor()
	queue := make(chan logFile)
	done := make(chan int)
	alert := make(chan int)
	numCheckers := len(config.Logfiles)

	// Start a checker for each log file listed in the config file.
	for i := 0; i < numCheckers; i++ {
		go checker(queue, done, alert)
	}

	// Populate the queue with the log files.
	for _, lf := range config.Logfiles {
		var lfile logFile
		lfile.Filename = lf
		queue <- lfile
	}

	// Wait for all the checkers to finish and listen for any alerts
	for i := 0; i < numCheckers; i++ {
		select {
		case <-done:
			continue
		case <-alert:
			alerted = true
		}
	}

	tools.Logger(monName, "Completed")
	return
}
