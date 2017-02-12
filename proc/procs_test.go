package proc

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnixProcPidRegex(t *testing.T) {
	assert.Equal(t, true, UnixProcPidRegex.MatchString("/proc/1234"))
	assert.Equal(t, false, UnixProcPidRegex.MatchString("/proc/foo"))
	assert.Equal(t, false, UnixProcPidRegex.MatchString("/proc/1234/bar"))
}

func TestIsUnixPidDir(t *testing.T) {
	p := filepath.Join("testdata", "proc", "1234")
	assert.Equal(t, true, isUnixPidDir(p))

	p = filepath.Join("testdata", "proc", "some-random-dir")
	assert.Equal(t, false, isUnixPidDir(p))

	p = filepath.Join("testdata", "proc", "something-that-does-not-exist")
	assert.Equal(t, false, isUnixPidDir(p))
}

func TestFindUnixProcesses(t *testing.T) {
	p := filepath.Join("testdata", "proc")
	procs := FindUnixProcesses(p)
	assert.Equal(t, 2, len(procs))

	p = filepath.Join("testdata", "does", "not", "exist")
	procs = FindUnixProcesses(p)
	assert.Equal(t, 0, len(procs))
}
