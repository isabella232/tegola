package main

import (
	"fmt"
	"os"

	"github.com/airmap/tegola/cmd/tegola/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
