package cfd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func RunCmd(cmdStr string) string {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		if err.Error() == "exit status 1" && stderr.String() == "" {
			return ""
		} else {
			log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
		}
	}

	return out.String()
}

func ParseAndFoundLineByLine(reg string, cmd_return string) bool {
	lines := strings.Split(cmd_return, "\n")

	for _, line := range lines {
		match, _ := regexp.MatchString(reg, line)

		if match {
			return match
		}
	}

	return false
}
