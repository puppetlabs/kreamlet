package main

import (
	"fmt"
	"os"

	"github.com/scotty-c/kream-v2/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
