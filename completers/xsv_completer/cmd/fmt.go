package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format CSV output (change field delimiter)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("ascii", false, "Use ASCII field and record separators.")
	fmtCmd.Flags().Bool("crlf", false, "Use '\r\n' line endings in the output.")
	fmtCmd.Flags().String("escape", "", "The escape character to use. When not specified,")
	fmtCmd.Flags().StringP("out-delimiter", "t", "", "The field delimiter for writing CSV data.")
	fmtCmd.Flags().String("quote", "", "The quote character to use. [default: \"]")
	fmtCmd.Flags().Bool("quote-always", false, "Put quotes around every value.")
	common.AddDelimiterFlag(fmtCmd)
	common.AddOutputFlag(fmtCmd)
	rootCmd.AddCommand(fmtCmd)
}
