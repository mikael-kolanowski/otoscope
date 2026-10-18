package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otoscope <port>...",
	Short: "Inspect which processes are listening on TCP ports",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			port, _ := strconv.Atoi(arg)
			fmt.Printf("checking %d\n", port)
			out, _ := GetTheStuff(port)
			fmt.Println(out)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
