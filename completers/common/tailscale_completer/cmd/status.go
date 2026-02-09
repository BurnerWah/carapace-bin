package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show state of tailscaled and its connections",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().Bool("active", false, "filter output to only peers with active sessions")
	statusCmd.Flags().Bool("browser", true, "Open a browser in web mode")
	statusCmd.Flags().Bool("header", false, "show column headers in table format")
	statusCmd.Flags().Bool("json", false, "output in JSON format")
	statusCmd.Flags().String("listen", "127.0.0.1:8384", "listen address for web mode; use port 0 for automatic")
	statusCmd.Flags().Bool("peers", true, "show status of peers")
	statusCmd.Flags().Bool("self", true, "show status of local machine")
	statusCmd.Flags().Bool("web", false, "run webserver with HTML showing status")

	rootCmd.AddCommand(statusCmd)
}
