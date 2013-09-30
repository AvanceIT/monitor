/*
Package logmon monitors the provided log files for errors

*/
package logmon

import (
	"bufio"
	"github.com/AvanceIT/monitor/tools"
	"os"
	"strings"
	"time"
)

const (
	numCheckers = 4 // Number of log files to check simultaneously
)

var monName string = "logmon"

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

func (lf logFile) checkFile() {
	lfile, err := os.Open(lf.Filename)
	if err != nil {
		tools.Logger(monName, "Unable to open logfile : "+lf.Filename)
		return
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
}

func checker(queue <-chan *logFile, done chan<- *logFile) {
	for lfile := range queue {
		message := "Checking " + lfile.Filename
		tools.Logger(monName, message)
		lfile.checkFile()
		time.Sleep(100 * time.Nanosecond)
		done <- lfile
	}
}

func RunChecks() (alerted bool) {
	tools.Logger(monName, "Starting")
	config := configMonitor()
	queue, done := make(chan *logFile), make(chan *logFile)

	for i := 0; i < numCheckers; i++ {
		go checker(queue, done)
	}

	for _, lf := range config.Logfiles {
		var lfile logFile
		lfile.Filename = lf
		queue <- &lfile
	}
	<-done

	tools.Logger(monName, "Completed")
	return
}
