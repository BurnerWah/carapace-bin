package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Compute basic statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().Bool("cardinality", false, "Show the cardinality.")
	statsCmd.Flags().Bool("everything", false, "Show all statistics available.")
	statsCmd.Flags().StringP("jobs", "j", "", "The number of jobs to run in parallel.")
	statsCmd.Flags().Bool("median", false, "Show the median.")
	statsCmd.Flags().Bool("mode", false, "Show the mode.")
	statsCmd.Flags().Bool("nulls", false, "Include NULLs in the population size for computing")
	statsCmd.Flags().StringP("select", "s", "", "Select a subset of columns to compute stats for.")
	common.AddDelimiterFlag(statsCmd)
	common.AddNoHeadersFlag(statsCmd)
	common.AddOutputFlag(statsCmd)
	rootCmd.AddCommand(statsCmd)
}
