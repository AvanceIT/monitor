/*
Package discmon provides the monitor to check current filesystem levels.

A configuration file defines the filesystems to check and the warning and
critical levels. It also has an ignore flag should the filesystem usage
be due to a known situation.
*/
package discmon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

// Type DiscConfig contains the relevant information for a filesystem.
type DiscConfig struct {
	FilesytemName string
	Ignore        bool
	Warn          int
	Crit          int
}

type Filesystems struct {
	DiscConfigs []DiscConfig
}

// Type FileSystemInfo contains the current usage of a filesystem.
type FileSystemInfo struct {
	FilesystemName string
	PercentUsed    int
}

// configMonitor reads the given configuration file and populates a
// Filesystems struct with the required details.
//
// The config file format is expect to be:-
// /some/filesystem::T::60::80
//
// which is the filesystem path followed by the ignore flag (T or F)
// and then the warn and critical percentages.
func configMonitor(fileName string) Filesystems {
	thisDiscConfig := DiscConfig{}
	fileSystems := Filsystems{}

	configFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("discmon: Error opening config file:\n\t%v\n", err)
	}
	defer configFile.Close()

	thisScanner := bufio.NewScanner(configFile)
	for thisScanner.Scan() {
		lineSplit := strings.Split(thisScanner.Text(), "::")
		thisDiscConfig.FilesystemName = lineSplit[0]
		if lineSplit[1] == "T" {
			thisDiscConfig.Ignore = true
		}
		thisDiscConfig.Warn = int(lineSplit[2])
		thisDiscConfig.Crit = int(lineSplit[3])
		fileSystems.DiscConfigs = append(fileSystems.DiscConfigs, thisDiscConfig)
	}

	return fileSystems
}

// getFsInfo returns the current information about a given filesystem.
func getFsInfo(fileSystem string) FileSystemInfo {
	var thisFstatFS syscall.Statfs_t
	thisFsInfo := FileSystemInfo{FileSystemName: fileSystem}

	thisFile, err := os.Open(fileSystem)
	if err != nil {
		fmt.Printf("discmon: error opening filesystem:\n\t%v\n", err)
	}
	defer thisFile.Close()

	thisFd := thisFile.Fd()
	syscall.Fstatfs(&thisFd, &thisFstatFS)

	thisPercentUsed = ((thisFstatFS.Blocks - thisFstatFS.Bavail) / thisFstatFS.Blocks) * 100
	thisFsInfo.PercentUsed = int(thisPercentUsed)

	return thisFsInfo
}

// RunChecks performs the checks required by this monitor. It returns
// a boolean value to denote whether an alert has been raised.
func RunChecks() bool {

	return alertRaised
}
