package utils

import (
	"bytes"
	"encoding/csv"
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
		if (err.Error() == "exit status 1") && stderr.String() == "" {
			return ""
		} else if err.Error() == "exit status 3" {
			// Not installed the command, mostly not supported the functions
			return ""
		} else {
			log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
		}
	}

	return out.String()
}

func ParseCSV(cmd_return string) [][]string {
	var returnValues [][]string = nil

	reader := csv.NewReader(strings.NewReader(cmd_return))
	returnValues, err := reader.ReadAll()

	if err != nil {
		return nil
	}

	return returnValues

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

func ReturnAndFoundLineByLine(reg string, cmd_return string) string {
	lines := strings.Split(cmd_return, "\n")

	for _, line := range lines {
		match, _ := regexp.MatchString(reg, line)

		if match {
			return line
		}
	}

	return ""
}

func ReturnSubValueOfFoundLineByLine(reg string, cmd_return string) [][]string {
	var returnValues [][]string = nil
	r := regexp.MustCompile(reg)

	lines := strings.Split(cmd_return, "\n")

	for _, line := range lines {
		match, _ := regexp.MatchString(reg, line)

		if match {
			oneNicInfo := r.FindStringSubmatch(line)
			if len(oneNicInfo) > 1 {
				returnValues = append(returnValues, oneNicInfo)
			}
		}
	}

	return returnValues
}
