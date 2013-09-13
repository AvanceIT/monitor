package fsmon

import (
	"testing"
)

func TestGetFsInfo(t *testing.T) {
	_ = getFsInfo("/")
}

func TestConfigMonitor(t *testing.T) {
	fs := configMonitor()
	fsLen := len(fs.FsConfigs)
	if fsLen == 0 {
		t.Fail()
	}
}
