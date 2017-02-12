// Package proc provides an API for handling processes.
package proc

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// UnixProcfs is the default procfs used for finding running processes on Unix.
const UnixProcfs = "/proc"

// UnixProcPidRegex is a compiled regex used for matching pid named directories.
var UnixProcPidRegex = regexp.MustCompile(`.*\/\d+$`)

// isUnixPidDir checks and returns true if the given directory meets the
// necessary conditions for finding a processes in a UNIX environment.
// `procPath`` should be the full path, eg: /proc/1234
func isUnixPidDir(procPath string) bool {
	dinfo, err := os.Stat(procPath)
	if err != nil {
		log.Error(err)
		return false
	}

	if dinfo.IsDir() && UnixProcPidRegex.MatchString(procPath) {
		// We are expecting in Unix that the pid dir has a cmdline
		cmdline := filepath.Join(procPath, "cmdline")

		finfo, err := os.Stat(cmdline)
		// If the proc directory is missing a cmdline, then
		// concider this to not be a valid proc directory.
		if err != nil || os.IsNotExist(err) {
			log.Warnf("Error getting stat on %q: %s", cmdline, err)
			return false
		}

		// We are expecting cmdline to be a file and not a directory.
		if finfo.IsDir() {
			log.Warnf("Expecting %q to be a file, but got a directory", cmdline)
			return false
		}

		return true
	}

	log.Warnf("Seems %q is not a pid directory", procPath)
	return false
}

// FindUnixProcesses returns a slice of running UNIX processes.
func FindUnixProcesses(procfs string) []*os.Process {
	var procs = []*os.Process{}
	if procfs == "" {
		procfs = UnixProcfs
	}

	items, err := ioutil.ReadDir(procfs)
	if err != nil {
		log.Errorf("Error getting content for path %s: %s", procfs, err)
	}

	for _, item := range items {
		fullpath := path.Join(procfs, item.Name())
		if !isUnixPidDir(fullpath) {
			log.Warnf("%q does not appear to be a pid directory.", fullpath)
			continue
		}
		pid, _ := strconv.Atoi(item.Name())
		proc, err := os.FindProcess(pid)
		if err != nil {
			log.Warnf("Could not create process with pid %d: %s", pid, err)
			continue
		}

		procs = append(procs, proc)
	}

	return procs
}
