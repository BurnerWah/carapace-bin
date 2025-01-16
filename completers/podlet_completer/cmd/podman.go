package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podmanCmd = &cobra.Command{
	Use:   "podman",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podmanCmd).Standalone()

	podmanCmd.Flags().String("cgroup-manager", "", "")
	podmanCmd.Flags().String("config", "", "")
	podmanCmd.Flags().String("conmon", "", "")
	podmanCmd.Flags().String("connection", "", "")
	podmanCmd.Flags().String("events-backend", "", "")
	podmanCmd.Flags().String("hooks-dir", "", "")
	podmanCmd.Flags().String("identity", "", "")
	podmanCmd.Flags().String("imagestore", "", "")
	podmanCmd.Flags().String("log-level", "", "")
	podmanCmd.Flags().String("module", "", "")
	podmanCmd.Flags().String("network-cmd-path", "", "")
	podmanCmd.Flags().String("network-config-dir", "", "")
	podmanCmd.Flags().String("out", "", "")
	podmanCmd.Flags().StringP("remote", "r", "", "")
	podmanCmd.Flags().String("root", "", "")
	podmanCmd.Flags().String("runroot", "", "")
	podmanCmd.Flags().String("runtime", "", "")
	podmanCmd.Flags().String("runtime-flag", "", "")
	podmanCmd.Flags().String("ssh", "", "")
	podmanCmd.Flags().String("storage-driver", "", "")
	podmanCmd.Flags().String("storage-opt", "", "")
	podmanCmd.Flags().Bool("syslog", false, "")
	podmanCmd.Flags().String("tmpdir", "", "")
	podmanCmd.Flags().String("transient-store", "", "")
	podmanCmd.Flags().String("url", "", "")
	podmanCmd.Flags().String("volumepath", "", "")
	rootCmd.AddCommand(podmanCmd)
}
