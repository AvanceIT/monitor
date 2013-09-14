/*
Package httpmon - Monitors a list of websites every 60secs

A connfiguration file in /etc/monitors/httpmon.cfg defines the sites to be monitored, one on each line.
*/

package httpmon

// Import ther packages we need ...
import (
	"github.com/AvanceIT/monitor/tools"
	"fmt"
)

var URL string
var monName string = "httpmon"
var alertRaised bool

/*
type URLlist struct {
	URL string
}
*/


func RunChecks() bool{
// Grab the list of URLs that we are to monitor ...

cl := tools.ReadConfig(monName)

for _, l := range cl {
	URL = l.Fields[0]
	fmt.Println(URL)
	//alertRaised = false


}
return alertRaised
}




