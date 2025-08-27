// Package util provides utility functions for output formatting.
package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tidwall/pretty"
)

// PrintJSON prints a JSON representation of the given value.
func PrintJSON(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%s\n", pretty.Color(pretty.Pretty(data), nil))
}

var defaultTableStyle = table.Style{
	Name: "CustomStyle",
	Box: table.BoxStyle{
		BottomLeft:       "└",
		BottomRight:      "┘",
		BottomSeparator:  "┴",
		Left:             "│",
		LeftSeparator:    "├",
		MiddleHorizontal: "─",
		MiddleSeparator:  "┼",
		MiddleVertical:   "│",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            "│",
		RightSeparator:   "┤",
		TopLeft:          "┌",
		TopRight:         "┐",
		TopSeparator:     "┬",
		UnfinishedRow:    "│",
	},
	Options: table.Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    true,
	},
}

// CreateTable creates a new table writer with the default style.
func CreateTable() table.Writer {
	t := table.NewWriter()
	t.SetStyle(defaultTableStyle)
	t.SetOutputMirror(os.Stdout)
	return t
}

// CreateHyperlink creates a clickable hyperlink for supported terminals.
func CreateHyperlink(url, text string) string {
	return fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, text)
}
