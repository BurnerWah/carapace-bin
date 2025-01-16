package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podman_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podman_volumeCmd).Standalone()

	podman_volumeCmd.Flags().String("cgroup-manager", "", "")
	podman_volumeCmd.Flags().String("config", "", "")
	podman_volumeCmd.Flags().String("conmon", "", "")
	podman_volumeCmd.Flags().String("connection", "", "")
	podman_volumeCmd.Flags().String("events-backend", "", "")
	podman_volumeCmd.Flags().String("hooks-dir", "", "")
	podman_volumeCmd.Flags().String("identity", "", "")
	podman_volumeCmd.Flags().String("imagestore", "", "")
	podman_volumeCmd.Flags().String("log-level", "", "")
	podman_volumeCmd.Flags().String("module", "", "")
	podman_volumeCmd.Flags().String("network-cmd-path", "", "")
	podman_volumeCmd.Flags().String("network-config-dir", "", "")
	podman_volumeCmd.Flags().String("out", "", "")
	podman_volumeCmd.Flags().StringP("remote", "r", "", "")
	podman_volumeCmd.Flags().String("root", "", "")
	podman_volumeCmd.Flags().String("runroot", "", "")
	podman_volumeCmd.Flags().String("runtime", "", "")
	podman_volumeCmd.Flags().String("runtime-flag", "", "")
	podman_volumeCmd.Flags().String("ssh", "", "")
	podman_volumeCmd.Flags().String("storage-driver", "", "")
	podman_volumeCmd.Flags().String("storage-opt", "", "")
	podman_volumeCmd.Flags().Bool("syslog", false, "")
	podman_volumeCmd.Flags().String("tmpdir", "", "")
	podman_volumeCmd.Flags().String("transient-store", "", "")
	podman_volumeCmd.Flags().String("url", "", "")
	podman_volumeCmd.Flags().String("volumepath", "", "")
	podmanCmd.AddCommand(podman_volumeCmd)
}
