package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Ports []uint
}

func parseArgs() (Config, error) {
	args := os.Args[1:]
	ports := make([]uint, 0, 1)
	for _, arg := range args {
		p, _ := strconv.Atoi(arg)
		ports = append(ports, uint(p))
	}

	return Config{Ports: ports}, nil
}

func main() {
	config, _ := parseArgs()
	for _, port := range config.Ports {
		fmt.Printf("checking %d\n", port)
	}
}
