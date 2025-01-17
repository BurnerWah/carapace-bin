package cmd

import (
	"github.com/carapace-sh/carapace"
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
	statsCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	statsCmd.Flags().Bool("everything", false, "Show all statistics available.")
	statsCmd.Flags().BoolP("help", "h", false, "Display this message")
	statsCmd.Flags().StringP("jobs", "j", "", "The number of jobs to run in parallel.")
	statsCmd.Flags().Bool("median", false, "Show the median.")
	statsCmd.Flags().Bool("mode", false, "Show the mode.")
	statsCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will NOT be interpreted")
	statsCmd.Flags().Bool("nulls", false, "Include NULLs in the population size for computing")
	statsCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	statsCmd.Flags().StringP("select", "s", "", "Select a subset of columns to compute stats for.")
	rootCmd.AddCommand(statsCmd)
}
