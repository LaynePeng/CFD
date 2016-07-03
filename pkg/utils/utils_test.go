package utils

import (
	"fmt"
	"strings"
	"testing"
)

var (
	content = "icp_dev0 - type=dh895xcc, inst_id=0, node_id=0, bdf=83:00:0, #accel=6, #engines=12, state=up"
)

func TestRunCmd(t *testing.T) {
	cmd := fmt.Sprintf("echo '%s'", content)
	ret := strings.TrimSpace(RunCmd(cmd)) // RunCMD will return with "\n"

	t.Log(ret)

	if ret != content {
		error_msg := fmt.Sprintf("\nRet: %s \nOri: %s", ret, content)
		t.Error(error_msg)
	}
}

func TestParseAndFoundLineByLine(t *testing.T) {
	cmd := fmt.Sprintf("echo '%s'", content)
	ret := RunCmd(cmd)
	regexp := "^(.+) - type=(\\S+), inst_id=0, node_id=(\\d+), bdf=.+, #accel=\\d+, #engines=(\\d+), state=(\\w+)"

	if !ParseAndFoundLineByLine(regexp, ret) {
		t.Error("Should be true!")
	}
}

func TestReturnAndFoundLineByLine(t *testing.T) {
	cmd := fmt.Sprintf("echo '%s'", content)
	ret := RunCmd(cmd)
	regexp := "^(.+) - type=(\\S+), inst_id=0, node_id=(\\d+), bdf=.+, #accel=\\d+, #engines=(\\d+), state=(\\w+)"

	match_line := ReturnAndFoundLineByLine(regexp, ret)

	t.Log(match_line)

	if match_line != content {
		t.Fail()
	}
}
