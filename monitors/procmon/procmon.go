// Package procmon monitors a provided list of processes are running
package procmon

import (
	"github.com/AvanceIT/monitor/tools"
	"os"
	"os/exec"
	"strings"
)

var monName string = "procmon"

type processes struct {
	ProcessName  string
	ProcessOwner string
}

type configOptions struct {
	HostName  string
	Processes []processes
}

// configMonitor reads the configuration file and sets up
// the configOptions structure. The configuration file just
// contains the process and process owner information to be
// monitored in the following format: -
//
// process::owner
func configMonitor() configOptions {
	HostName, _ := os.Hostname()
	configuration := configOptions{HostName: HostName}
	thisProcess := processes{}

	cl := tools.ReadConfig(monName)
	for _, l := range cl {
		thisProcess.ProcessName = l.Fields[0]
		thisProcess.ProcessOwner = l.Fields[1]
		configuration.Processes = append(configuration.Processes, thisProcess)
	}

	return configuration
}

// getProcessList grabs the current running processes using ps
// and returns them as a []string
func getProcessList() []string {
	var thisLine []byte
	var psLines []string

	ps := exec.Command("ps", "-e", "-ouser=,pid=,ppid=,comm=")
	output, _ := ps.Output()
	for _, thisChar := range output {
		if string(thisChar) != "\n" {
			thisLine = append(thisLine, thisChar)
		} else {
			psLines = append(psLines, string(thisLine))
			thisLine = []byte{}
		}
	}

	return psLines
}

// RunChecks performs the checks required by this monitor. It returns
// boolean value to denote whether an alert has been raised.
func RunChecks() bool {
	tools.Logger(monName, "Starting")
	processList := getProcessList()
	var runningProcesses []processes
	var processFound bool
	var alertRaised bool
	var alertString string
	thisConfig := configMonitor()

	for _, processListLine := range processList {
		processListFields := strings.Fields(processListLine)
		thisProcesses := processes{
			ProcessName:  processListFields[3],
			ProcessOwner: processListFields[0],
		}
		runningProcesses = append(runningProcesses, thisProcesses)
	}

	for _, checkProcesses := range thisConfig.Processes {
		tools.Logger(monName, "checking " + checkProcesses.ProcessName)
		for _, eachRunningProcess := range runningProcesses {
			if checkProcesses.ProcessName == eachRunningProcess.ProcessName {
				if checkProcesses.ProcessOwner == eachRunningProcess.ProcessOwner {
					processFound = true
				}
			}
		}
		if !processFound {
			alertString = ("ProcMon: " + checkProcesses.ProcessName + " not running for user " + checkProcesses.ProcessOwner)
			tools.RaiseAlert(alertString, 99)
			alertRaised = true
		}
		processFound = false
	}

	tools.Logger(monName, "Completed")
	return alertRaised
}
