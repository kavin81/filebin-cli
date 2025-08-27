// Package main ...
package main

import (
	"fmt"
	"os"

	"filebin-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
