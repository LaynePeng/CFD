package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func main() {
	cmd := exec.Command("cfd", "t", "-f", "gpu")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	value, err := strconv.ParseBool(out.String())

	if value {
		fmt.Println("It is true!")
	} else {
		fmt.Println("It is false!")
	}
}
