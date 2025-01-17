package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("ascii", false, "Use ASCII field and record separators.")
	fmtCmd.Flags().Bool("crlf", false, "Use '\r\n' line endings in the output.")
	fmtCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	fmtCmd.Flags().String("escape", "", "The escape character to use. When not specified,")
	fmtCmd.Flags().BoolP("help", "h", false, "Display this message")
	fmtCmd.Flags().StringP("out-delimiter", "t", "", "The field delimiter for writing CSV data.")
	fmtCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	fmtCmd.Flags().String("quote", "", "The quote character to use. [default: \"]")
	fmtCmd.Flags().Bool("quote-always", false, "Put quotes around every value.")
	rootCmd.AddCommand(fmtCmd)
}
