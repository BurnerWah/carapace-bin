package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server for controlling Tailscale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webCmd).Standalone()

	webCmd.Flags().Bool("cgi", false, "run as CGI script")
	webCmd.Flags().String("listen", "localhost:8088", "listen address; use port 0 for automatic")
	webCmd.Flags().String("origin", "", "origin at which the web UI is served")
	webCmd.Flags().String("prefix", "", "URL prefix added to requests")
	webCmd.Flags().Bool("readonly", false, "")

	rootCmd.AddCommand(webCmd)

	// TODO: Add flag completion
}
