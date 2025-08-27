package cmd

import (
	"fmt"
	"strings"

	"filebin-cli/api"
	"filebin-cli/errors"
	"filebin-cli/types"
	"filebin-cli/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete <bin>[/<file>]",
	Short: "Remove bin or file from bin",
	Long: `Remove a bin or specific file from a filebin. The bin will be permanently deleted if no file is specified.

Examples:
  filebin delete team-docs
  filebin delete temp-share/archive.zip
  filebin delete public-files presentation.pptx`,
	Args:    cobra.RangeArgs(1, 2),
	Aliases: []string{"remove", "del", "rm"},
	GroupID: "bin",
	RunE:    runRemoveCommand,
}

func runRemoveCommand(cmd *cobra.Command, args []string) error {
	resource := parseResourceFromArgs(args)
	ops := api.NewOperations()

	if verbose {
		ops.EnableDebug()
	}

	err := ops.DeleteResource(resource)
	if err != nil {
		return handleRemoveError(err, resource)
	}

	printRemoveSuccess(resource)
	return nil
}

func parseResourceFromArgs(args []string) types.Resource {
	if len(args) == 1 {
		if strings.Contains(args[0], "/") {
			parts := strings.SplitN(args[0], "/", 2)
			return types.NewFileResource(parts[0], parts[1])
		}
		return types.NewBinResource(args[0])
	}
	return types.NewFileResource(args[0], args[1])
}

func handleRemoveError(err error, resource types.Resource) error {
	switch e := err.(type) {
	case *errors.BinNotFoundError:
		binURL := fmt.Sprintf("https://filebin.net/%s", e.BinID)
		binLink := util.CreateHyperlink(binURL, e.BinID)
		fmt.Printf("%s Bin %s not found\n",
			color.HiRedString("error:"),
			color.New(color.Faint).Sprint(binLink))
	case *errors.FileNotFoundError:
		binURL := fmt.Sprintf("https://filebin.net/%s", e.BinID)
		fileURL := fmt.Sprintf("https://filebin.net/%s/%s", e.BinID, e.Filename)
		binLink := util.CreateHyperlink(binURL, e.BinID)
		fileLink := util.CreateHyperlink(fileURL, e.Filename)
		fmt.Printf("%s File %s not found in bin %s\n",
			color.HiRedString("error:"),
			color.New(color.Faint).Sprint(fileLink),
			color.New(color.Faint).Sprint(binLink))
	default:
		targetType := "bin"
		if resource.IsFile() {
			targetType = "file"
		}
		fmt.Printf("%s Failed to remove %s: %v\n",
			color.HiRedString("error:"),
			targetType,
			err)
	}
	return err
}

func printRemoveSuccess(resource types.Resource) {
	if resource.IsFile() {
		fileURL := fmt.Sprintf("https://filebin.net/%s/%s", resource.BinID, resource.Filename)
		binURL := fmt.Sprintf("https://filebin.net/%s", resource.BinID)
		fileLink := util.CreateHyperlink(fileURL, resource.Filename)
		binLink := util.CreateHyperlink(binURL, resource.BinID)
		fmt.Printf("%s File %s removed from bin %s\n",
			color.HiGreenString("success:"),
			color.New(color.Faint).Sprint(fileLink),
			color.New(color.Faint).Sprint(binLink))
	} else {
		binURL := fmt.Sprintf("https://filebin.net/%s", resource.BinID)
		binLink := util.CreateHyperlink(binURL, resource.BinID)
		fmt.Printf("%s Bin %s removed\n",
			color.HiGreenString("success:"),
			color.New(color.Faint).Sprint(binLink))
	}
}

func init() {
	rootCmd.AddCommand(deleteCommand)
}
