package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xsv_completer/cmd/common"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split CSV data into many files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().String("filename", "", "A filename template to use when constructing")
	splitCmd.Flags().StringP("jobs", "j", "", "The number of spliting jobs to run in parallel.")
	splitCmd.Flags().StringP("size", "s", "", "The number of records to write into each chunk.")
	common.AddDelimiterFlag(splitCmd)
	common.AddNoHeadersFlag(splitCmd)
	rootCmd.AddCommand(splitCmd)
}
