package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frequencyCmd = &cobra.Command{
	Use:   "frequency",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frequencyCmd).Standalone()

	frequencyCmd.Flags().BoolP("asc", "a", false, "Sort the frequency tables in ascending order by")
	frequencyCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	frequencyCmd.Flags().BoolP("help", "h", false, "Display this message")
	frequencyCmd.Flags().StringP("jobs", "j", "", "The number of jobs to run in parallel.")
	frequencyCmd.Flags().StringP("limit", "l", "", "Limit the frequency table to the N most common")
	frequencyCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will NOT be included")
	frequencyCmd.Flags().Bool("no-nulls", false, "Don't include NULLs in the frequency table.")
	frequencyCmd.Flags().StringP("output", "o", "", "Write output to <file> instead of stdout.")
	frequencyCmd.Flags().StringP("select", "s", "", "Select a subset of columns to compute frequencies")
	rootCmd.AddCommand(frequencyCmd)
}
