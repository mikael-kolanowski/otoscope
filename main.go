package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Ports []int
}

func parseArgs() (Config, error) {
	args := os.Args[1:]
	ports := make([]int, 0, 1)
	for _, arg := range args {
		p, _ := strconv.Atoi(arg)
		ports = append(ports, p)
	}

	return Config{Ports: ports}, nil
}

func main() {
	config, _ := parseArgs()
	for _, port := range config.Ports {
		fmt.Printf("checking %d\n", port)
		out, _ := GetTheStuff(port)
		fmt.Println(out)
	}
}
