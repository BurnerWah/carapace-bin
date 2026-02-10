package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View current serve configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_statusCmd).Standalone()

	serve_statusCmd.Flags().Bool("json,", false, "output JSON")

	serveCmd.AddCommand(serve_statusCmd)
}
