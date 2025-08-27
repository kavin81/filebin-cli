package cmd

import (
	"fmt"

	"filebin-cli/api"
	"filebin-cli/errors"
	"filebin-cli/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock <bin>",
	Short: "Lock bin to prevent uploads",
	Long: `Lock a bin to prevent further file uploads. This action is irreversible once applied.

Examples:
  filebin lock team-docs
  filebin lock temp-share
  filebin lock public-files`,
	Args:    cobra.ExactArgs(1),
	GroupID: "bin",
	RunE:    runLockCommand,
}

func runLockCommand(cmd *cobra.Command, args []string) error {
	binID := args[0]
	ops := api.NewOperations()

	if verbose {
		ops.EnableDebug()
	}

	err := ops.LockBin(binID)
	if err != nil {
		switch e := err.(type) {
		case *errors.BinNotFoundError:
			fmt.Printf("%s Bin %s not found\n",
				color.HiRedString("error:"),
				color.New(color.Faint).Sprint(e.BinID))
		default:
			fmt.Printf("%s Failed to lock bin: %v\n",
				color.HiRedString("error:"),
				err)
		}
		return err
	}

	printLockSuccess(binID)
	return nil
}

func printLockSuccess(binID string) {
	binURL := fmt.Sprintf("https://filebin.net/%s", binID)
	binLink := util.CreateHyperlink(binURL, binID)
	fmt.Printf("%s Bin %s locked\n",
		color.HiGreenString("success:"),
		color.New(color.Faint).Sprint(binLink))
}

func init() {
	rootCmd.AddCommand(lockCmd)
}
