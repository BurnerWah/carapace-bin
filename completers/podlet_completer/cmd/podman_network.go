package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podman_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podman_networkCmd).Standalone()

	podman_networkCmd.Flags().String("cgroup-manager", "", "")
	podman_networkCmd.Flags().String("config", "", "")
	podman_networkCmd.Flags().String("conmon", "", "")
	podman_networkCmd.Flags().String("connection", "", "")
	podman_networkCmd.Flags().String("events-backend", "", "")
	podman_networkCmd.Flags().String("hooks-dir", "", "")
	podman_networkCmd.Flags().String("identity", "", "")
	podman_networkCmd.Flags().String("imagestore", "", "")
	podman_networkCmd.Flags().String("log-level", "", "")
	podman_networkCmd.Flags().String("module", "", "")
	podman_networkCmd.Flags().String("network-cmd-path", "", "")
	podman_networkCmd.Flags().String("network-config-dir", "", "")
	podman_networkCmd.Flags().String("out", "", "")
	podman_networkCmd.Flags().StringP("remote", "r", "", "")
	podman_networkCmd.Flags().String("root", "", "")
	podman_networkCmd.Flags().String("runroot", "", "")
	podman_networkCmd.Flags().String("runtime", "", "")
	podman_networkCmd.Flags().String("runtime-flag", "", "")
	podman_networkCmd.Flags().String("ssh", "", "")
	podman_networkCmd.Flags().String("storage-driver", "", "")
	podman_networkCmd.Flags().String("storage-opt", "", "")
	podman_networkCmd.Flags().Bool("syslog", false, "")
	podman_networkCmd.Flags().String("tmpdir", "", "")
	podman_networkCmd.Flags().String("transient-store", "", "")
	podman_networkCmd.Flags().String("url", "", "")
	podman_networkCmd.Flags().String("volumepath", "", "")
	podmanCmd.AddCommand(podman_networkCmd)
}
