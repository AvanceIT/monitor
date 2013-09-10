/*
Package tools provides various functions for working with
the monitors and alerts in the monitoring client.
*/
package tools

import (
	"fmt"
	"monitor/xmltools"
	"os"
)

// getMonitorName takes the alertString and grabs the monitor
// name from the front. It then returns the monitorName and
// also the alertMessage without the monitorName at the start
func getMonitorName(alertMessage string) (string, string) {
	var monitorName []byte

	for messageStart, thisChar := range alertMessage {
		if string(thisChar) != ":" {
			monitorName = append(monitorName, byte(thisChar))
		} else {
			alertMessage = alertMessage[messageStart+2:]
			break
		}
	}

	return string(monitorName), alertMessage
}

// RaiseAlert formats the alertString into an XML message
// and passes it to the monitoring server
func RaiseAlert(alertMessage string, alertLevel int) {
	var monitorName string
	thisHostName, _ := os.Hostname()

	monitorName, alertMessage = getMonitorName(alertMessage)
	alertData := xmltools.MonResult{
		MonName:    monitorName,
		AlertLevel: alertLevel,
		HostName:   thisHostName,
		Detail:     alertMessage,
	}
	alertString := xmltools.CreateAlert(alertData)

	// TODO(JP): Implement client connection in RaiseAlert()
	fmt.Printf(alertString)
}
