// Package cmd ...
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "filebin",
}

func Execute() error {
	return rootCmd.Execute()
}

var (
	verbose    bool
	jsonOutput bool
	icons      bool
)

func init() {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "bin",
		Title: "Bin Commands",
	})

	rootCmd.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false, "Enable verbose output",
	)
	rootCmd.PersistentFlags().BoolVar(
		&jsonOutput, "json", false, "Enable JSON output",
	)
	rootCmd.PersistentFlags().BoolVar(
		&icons, "icons", false, "Set icon display mode (experimental)",
	)
}
