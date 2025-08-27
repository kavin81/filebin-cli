package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"filebin-cli/api"
	"filebin-cli/errors"
	"filebin-cli/types"
	"filebin-cli/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	clientID       string
	customFilename string
)

var pushCmd = &cobra.Command{
	Use:   "push <bin> <file>",
	Short: "Upload file to bin",
	Long: `Upload a file to a filebin for sharing. If the bin doesn't exist, it will be created automatically.

Examples:
  filebin push team-docs presentation.pptx
  filebin push temp-share ./archive.zip
  filebin push public-files image.jpg --filename "hero-image.jpg"`,
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"upload", "put"},
	GroupID: "bin",
	RunE:    runPushCommand,
}

func runPushCommand(cmd *cobra.Command, args []string) error {
	binID := args[0]
	filePath := args[1]

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("%s File %s does not exist\n",
			color.HiRedString("error:"),
			color.New(color.Faint).Sprint(filePath))
		return err
	}

	// Get absolute path for display
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		absPath = filePath
	}

	config := types.NewUploadConfig(binID, filePath)

	// Set custom filename if provided
	if customFilename != "" {
		config.SetCustomFilename(customFilename)
	}

	// Set client ID if provided
	if clientID != "" {
		config.SetClientID(clientID)
	}

	ops := api.NewOperations()
	if verbose {
		ops.EnableDebug()
	}

	err = ops.UploadFile(config)
	if err != nil {
		return handleUploadError(err, config)
	}

	printUploadSuccess(config, absPath)
	return nil
}

func handleUploadError(err error, config types.UploadConfig) error {
	switch e := err.(type) {
	case *errors.APIError:
		switch e.StatusCode {
		case 400:
			fmt.Printf("%s Invalid bin ID or filename: %s\n",
				color.HiRedString("error:"),
				e.Message)
		case 403:
			fmt.Printf("%s Storage limitation reached\n",
				color.HiRedString("error:"))
		case 405:
			fmt.Printf("%s Bin %s is locked and cannot be written to\n",
				color.HiRedString("error:"),
				color.New(color.Faint).Sprint(config.BinID))
		default:
			fmt.Printf("%s Upload failed: %s\n",
				color.HiRedString("error:"),
				e.Message)
		}
	default:
		fmt.Printf("%s Failed to upload file: %v\n",
			color.HiRedString("error:"),
			err)
	}
	return err
}

func printUploadSuccess(config types.UploadConfig, originalPath string) {
	binURL := fmt.Sprintf("https://filebin.net/%s", config.BinID)
	binLink := util.CreateHyperlink(binURL, config.BinID)

	fileURL := fmt.Sprintf("https://filebin.net/%s/%s", config.BinID, config.Filename)
	fileLink := util.CreateHyperlink(fileURL, config.Filename)

	fmt.Printf("%s File %s uploaded to bin %s as %s\n",
		color.HiGreenString("success:"),
		color.New(color.Faint).Sprint(originalPath),
		color.New(color.Faint).Sprint(binLink),
		color.New(color.Faint).Sprint(fileLink))
}

func init() {
	pushCmd.Flags().StringVar(&clientID, "client-id", "", "Custom client identifier")
	pushCmd.Flags().StringVar(&customFilename, "filename", "", "Custom filename for the uploaded file")
	rootCmd.AddCommand(pushCmd)
}
