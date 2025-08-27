package cmd

import (
	"fmt"

	"filebin-cli/api"
	"filebin-cli/types"
	"filebin-cli/util"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <bin>",
	Short: "Show information about bin",
	Long: `Display detailed information about a bin including files and metadata. Shows file sizes, upload dates, and bin status.

Examples:
  filebin show team-docs
  filebin show temp-share
  filebin show public-files --json`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"info"},
	GroupID: "bin",
	RunE:    runShowCommand,
}

func runShowCommand(cmd *cobra.Command, args []string) error {
	binID := args[0]
	ops := api.NewOperations()

	if verbose {
		ops.EnableDebug()
	}

	result, err := ops.GetBinInfo(binID)
	if err != nil {
		return fmt.Errorf("failed to fetch bin info: %w", err)
	}

	if jsonOutput {
		util.PrintJSON(result)
		return nil
	}

	displayBinInfo(result)
	return nil
}

func displayBinInfo(result *types.BinInfo) {
	binURL := fmt.Sprintf("https://filebin.net/%s", result.Bin.ID)
	binLink := util.CreateHyperlink(binURL, result.Bin.ID)

	t := util.CreateTable()
	t.SetTitle(fmt.Sprintf("Bin: %s", binLink))
	t.AppendHeader(table.Row{"Size", "Files", "Readonly", "Created", "Updated", "Expires"})

	readonly := color.GreenString("No")
	if result.Bin.Readonly {
		readonly = color.RedString("Yes")
	}

	t.AppendRow(table.Row{
		result.Bin.BytesReadable,
		result.Bin.Files,
		readonly,
		result.Bin.CreatedAtRelative,
		result.Bin.UpdatedAtRelative,
		result.Bin.ExpiredAtRelative,
	})
	t.Style().Title.Align = text.AlignCenter
	t.Render()

	if result.Bin.Files > 0 {
		displayFileList(result)
	}
}

func displayFileList(result *types.BinInfo) {
	fmt.Println()
	tf := util.CreateTable()
	tf.SetTitle("Files")
	tf.AppendHeader(table.Row{"Filename", "Type", "Size"})

	for _, f := range result.Files {
		fileURL := fmt.Sprintf("https://filebin.net/%s/%s", result.Bin.ID, f.Filename)
		fileLink := util.CreateHyperlink(fileURL, f.Filename)

		tf.AppendRow(table.Row{
			fileLink,
			f.ContentType,
			f.BytesReadable,
		})
	}
	tf.Style().Title.Align = text.AlignCenter
	tf.Render()
}

func init() {
	rootCmd.AddCommand(showCmd)
}
