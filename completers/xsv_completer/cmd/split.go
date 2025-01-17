package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split CSV data into many files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().StringP("delimiter", "d", "", "The field delimiter for reading CSV data.")
	splitCmd.Flags().String("filename", "", "A filename template to use when constructing")
	splitCmd.Flags().BoolP("help", "h", false, "Display this message")
	splitCmd.Flags().StringP("jobs", "j", "", "The number of spliting jobs to run in parallel.")
	splitCmd.Flags().BoolP("no-headers", "n", false, "When set, the first row will NOT be interpreted")
	splitCmd.Flags().StringP("size", "s", "", "The number of records to write into each chunk.")
	rootCmd.AddCommand(splitCmd)
}
