// Package xmltools provides simple tools for handling XML monitoring data
package xmltools

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// MonResult groups the data from the monitor script and uses
// field tags to define the XML tags.
type MonResult struct {
	HostName   string `xml:"HostName,attr"`
	MonName    string `xml:"MonName"`
	TimeRcvd   string `xml:"TimeRcvd"`
	TimeRptd   string `xml:"TimeRptd"`
	AlertLevel int    `xml:"AlertLevel"`
	Detail     string `xml:"Detail"`
}

var timeFormat string = "Mon Jan 02 2006 15:04:05 MST"

// Take a MonResult and format it to XML and then dump it to STDOUT
// to practice using XML encoding.
func DumpXML(data MonResult) {
	timeNow := time.Now()
	data.TimeRcvd = timeNow.Format(timeFormat)
	output, err := xml.MarshalIndent(&data, " ", "    ")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	os.Stdout.Write(output)
}

// CreateAlert takes a MonResult struct and returns it as a string
// suitable for passing to the alert server.
func CreateAlert(data MonResult) string {
	timeNow := time.Now()
	data.TimeRcvd = timeNow.Format(timeFormat)

	output, err := xml.MarshalIndent(&data, "", "    ")
	if err != nil {
		fmt.Printf("CreateAlert error: %v\n", err)
	}

	return string(output)
}
