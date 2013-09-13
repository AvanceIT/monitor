package fsmon

import (
	"testing"
)

func TestGetFsInfo(t *testing.T) {
	_ = getFsInfo("/")
}
