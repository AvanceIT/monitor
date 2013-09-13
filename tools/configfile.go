package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ConfigDir string = "/etc/monitor/"

type ConfigLine struct {
	Fields []string
}

// Function ReadConfig takes the monitor name and uses it to
// form a file name. It reads this file and splits the lines
// into fields and returns them in a []string
func ReadConfig(mn string) []ConfigLine {
	var lines []ConfigLine
	var fn string = ConfigDir + mn + ".cfg"
	var tl string
	var cl ConfigLine

	_, err := os.Stat(fn)
	if err != nil {
		return lines
	}

	file, err := os.Open(fn)
	if err != nil {
		fmt.Printf("\n%s: Cannot open %s:\n\t%v\n", mn, fn)
		return lines
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		tl = s.Text()
		if tl[0] == '#' {
			continue
		}
		cl.Fields = strings.Split(tl, "::")
		lines = append(lines, cl)
	}

	return lines
}
