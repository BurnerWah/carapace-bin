package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var frequencyCmd = &cobra.Command{
	Use:   "frequency",
	Short: "Show frequency tables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frequencyCmd).Standalone()

	frequencyCmd.Flags().BoolP("asc", "a", false, "Sort the frequency tables in ascending order by")
	frequencyCmd.Flags().StringP("jobs", "j", "", "The number of jobs to run in parallel.")
	frequencyCmd.Flags().StringP("limit", "l", "", "Limit the frequency table to the N most common")
	frequencyCmd.Flags().Bool("no-nulls", false, "Don't include NULLs in the frequency table.")
	frequencyCmd.Flags().StringP("select", "s", "", "Select a subset of columns to compute frequencies")
	common.AddDelimiterFlag(frequencyCmd)
	common.AddNoHeadersFlag(frequencyCmd)
	common.AddOutputFlag(frequencyCmd)
	rootCmd.AddCommand(frequencyCmd)
}
