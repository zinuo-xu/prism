package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/zinuo-xu/prism/internal/model"
)

var (
	green  = color.New(color.FgGreen)
	red    = color.New(color.FgRed)
	yellow = color.New(color.FgYellow)
)

func Terminal(deltas []model.Delta) {
	if len(deltas) == 0 {
		green.Println("No differences found!")
		return
	}
	for _, d := range deltas {
		switch d.Type {
		case model.Added:
			fmt.Printf("  %s %s: %s\n", green.Sprint("+"), d.Path, d.Message)
		case model.Removed:
			fmt.Printf("  %s %s: %s\n", red.Sprint("-"), d.Path, d.Message)
		case model.Changed:
			fmt.Printf("  %s %s: %s\n", yellow.Sprint("~"), d.Path, d.Message)
		case model.Reordered:
			fmt.Printf("  %s %s: %s\n", yellow.Sprint("<>"), d.Path, d.Message)
		case model.TypeChanged:
			fmt.Printf("  %s %s: %s\n", red.Sprint("T"), d.Path, d.Message)
		}
	}
}
