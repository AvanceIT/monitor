package tools

import (
	"fmt"
	"os"
	"time"
)

var logfile string = "/tmp/monitor.log"
var timeFormat string = "Jan 02 15:04:06"

func openLogfile() *os.File {
	var lf *os.File

	_, err := os.Stat(logfile)
	if err != nil {
		lf, err = os.Create(logfile)
	} else {
		lf, err = os.Open(logfile)
	}

	if err != nil {
		fmt.Printf("Monitor could not open log file: %v\n", err)
	}

	return lf
}

// Function Logger takes monitor name and message as strings and writes them
// to a logfile.
func Logger(mn string, msg string) {
	lf := openLogfile()
	defer lf.Close()

	tn := time.Now()
	ts := tn.Format(timeFormat)
	hn, _ := os.Hostname()
	e := ts + " " + hn + " " + mn + ": " + msg + "\n"
	lf.WriteString(e)
}
