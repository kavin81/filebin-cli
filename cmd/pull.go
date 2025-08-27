package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"filebin-cli/api"
	"filebin-cli/errors"
	"filebin-cli/types"
	"filebin-cli/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var outputPath string

var pullCmd = &cobra.Command{
	Use:   "pull <bin>[/<file>]",
	Short: "Download bin or file from bin",
	Long: `Download files from a filebin. Downloads entire bin as archive if no file is specified.

Examples:
  filebin pull team-docs
  filebin pull temp-share/archive.zip
  filebin pull public-files presentation.pptx --output ./downloads`,
	Args:    cobra.RangeArgs(1, 2),
	Aliases: []string{"get", "fetch"},
	GroupID: "bin",
	RunE:    runGetCommand,
}

func runGetCommand(cmd *cobra.Command, args []string) error {
	resource := parseDownloadResource(args)

	if outputPath == "" {
		var err error
		outputPath, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}
	}

	config := types.DownloadConfig{
		OutputPath: outputPath,
		Resource:   resource,
	}

	ops := api.NewOperations()
	if verbose {
		ops.EnableDebug()
	}

	err := ops.DownloadResource(config)
	if err != nil {
		return handleDownloadError(err, resource)
	}

	printDownloadSuccess(config)
	return nil
}

func parseDownloadResource(args []string) types.Resource {
	if len(args) == 1 {
		if strings.Contains(args[0], "/") {
			parts := strings.SplitN(args[0], "/", 2)
			return types.NewFileResource(parts[0], parts[1])
		}
		return types.NewBinResource(args[0])
	}
	return types.NewFileResource(args[0], args[1])
}

func handleDownloadError(err error, resource types.Resource) error {
	switch e := err.(type) {
	case *errors.BinNotFoundError:
		fmt.Printf("%s Bin %s not found\n",
			color.HiRedString("error:"),
			color.New(color.Faint).Sprint(e.BinID))
	case *errors.FileNotFoundError:
		fmt.Printf("%s File %s not found in bin %s\n",
			color.HiRedString("error:"),
			color.New(color.Faint).Sprint(e.Filename),
			color.New(color.Faint).Sprint(e.BinID))
	default:
		targetType := "bin"
		if resource.IsFile() {
			targetType = "file"
		}
		fmt.Printf("%s Failed to download %s: %v\n",
			color.HiRedString("error:"),
			targetType,
			err)
	}
	return err
}

func printDownloadSuccess(config types.DownloadConfig) {
	outputFile := config.GetOutputPath()
	absPath, _ := filepath.Abs(outputFile)

	if config.Resource.IsFile() {
		fmt.Printf("%s File %s downloaded to %s\n",
			color.HiGreenString("success:"),
			color.New(color.Faint).Sprint(config.Resource.Filename),
			color.New(color.Faint).Sprint(absPath))
	} else {
		binID := config.Resource.BinID
		fileURL := fmt.Sprintf("https://filebin.net/%s/%s", binID, config.Resource.Filename)
		fileLink := util.CreateHyperlink(fileURL, config.Resource.Filename)
		fmt.Printf("%s Downloaded %s to %s\n",
			color.HiGreenString("success:"),
			color.New(color.Faint).Sprint(fileLink),
			absPath)
	}
}

func init() {
	pullCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output directory (default: current directory)")
	rootCmd.AddCommand(pullCmd)
}
