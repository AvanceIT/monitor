/*
Package tools provides various functions for working with
the monitors and alerts in the monitoring client.
*/
package tools

import (
	"encoding/gob"
	"fmt"
	"github.com/AvanceIT/monitor/xmltools"
	"net"
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

// RaiseAlert formats the alertString into an Gob encoded message
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
	srv, err := net.Dial("tcp", "192.168.0.5:2468")
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		return
	}
	defer srv.Close()
	enc := gob.NewEncoder(srv)
	err = enc.Encode(&alertData)
	if err != nil {
		fmt.Printf("encode error: %v\n", err)
	}
}
