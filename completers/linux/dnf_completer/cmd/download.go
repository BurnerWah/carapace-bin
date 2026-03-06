package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Short:   "Download software to the current directory",
	GroupID: "commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()

	downloadCmd.Flags().Bool("alldeps", false, "When running with --resolve, download all dependencies (do not exclude already installed ones)")
	downloadCmd.Flags().StringSlice("arch", nil, "Limit to packages of given architectures.")
	downloadCmd.Flags().String("destdir", "", "Set directory used for downloading packages to")
	downloadCmd.Flags().Bool("resolve", false, "Resolve and download needed dependencies")
	downloadCmd.Flags().Bool("source", false, "Alias for '--srpm'")
	downloadCmd.Flags().Bool("srpm", false, "Download the src.rpm instead")
	downloadCmd.Flags().Bool("url", false, "Print a URL where the rpms can be downloaded instead of downloading")
	downloadCmd.Flags().StringSlice("urlprotocol", nil, "When running with --url, limit to specific protocols")
	downloadCmd.Flag("source").Hidden = true
	rootCmd.AddCommand(downloadCmd)
}
