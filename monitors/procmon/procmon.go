// Package procmon monitors a provided list of processes are running
package procmon

import (
	"bufio"
	"fmt"
	"github.com/AvanceIT/monitor/tools"
	"os"
	"os/exec"
	"strings"
)

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
func configMonitor(fileName string) configOptions {
	HostName, _ := os.Hostname()
	configuration := configOptions{HostName: HostName}
	thisProcesses := processes{}
	var thisLine string

	configFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer configFile.Close()

	thisScanner := bufio.NewScanner(configFile)
	for thisScanner.Scan() {
		thisLine = thisScanner.Text()
		if thisLine[0] == '#' {
			continue
		}
		lineSplit := strings.Split(thisLine, "::")
		thisProcesses.ProcessName = lineSplit[0]
		thisProcesses.ProcessOwner = lineSplit[1]
		configuration.Processes = append(configuration.Processes, thisProcesses)
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
	processList := getProcessList()
	var runningProcesses []processes
	var processFound bool
	var alertRaised bool
	var alertString string
	thisConfig := configMonitor("/tmp/procmon.cfg")

	for _, processListLine := range processList {
		processListFields := strings.Fields(processListLine)
		thisProcesses := processes{
			ProcessName:  processListFields[3],
			ProcessOwner: processListFields[0],
		}
		runningProcesses = append(runningProcesses, thisProcesses)
	}

	for _, checkProcesses := range thisConfig.Processes {
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

	return alertRaised
}
