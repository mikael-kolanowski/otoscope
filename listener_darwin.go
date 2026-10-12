package main

import (
	"fmt"
	"os/exec"
)

func prepareLsofCommand(port int) []string {
	return []string{"lsof", "-nP", fmt.Sprintf("-iTCP:%d", port), "-sTCP:LISTEN"}
}

func GetTheStuff(port int) (string, error) {
	lsof := prepareLsofCommand(port)
	cmd := exec.Command("lsof", lsof[1:]...)

	outBytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(outBytes), nil
}
