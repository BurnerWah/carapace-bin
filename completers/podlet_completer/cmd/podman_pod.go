package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podman_podCmd = &cobra.Command{
	Use:   "pod",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podman_podCmd).Standalone()

	podman_podCmd.Flags().String("cgroup-manager", "", "")
	podman_podCmd.Flags().String("config", "", "")
	podman_podCmd.Flags().String("conmon", "", "")
	podman_podCmd.Flags().String("connection", "", "")
	podman_podCmd.Flags().String("events-backend", "", "")
	podman_podCmd.Flags().String("hooks-dir", "", "")
	podman_podCmd.Flags().String("identity", "", "")
	podman_podCmd.Flags().String("imagestore", "", "")
	podman_podCmd.Flags().String("log-level", "", "")
	podman_podCmd.Flags().String("module", "", "")
	podman_podCmd.Flags().String("network-cmd-path", "", "")
	podman_podCmd.Flags().String("network-config-dir", "", "")
	podman_podCmd.Flags().String("out", "", "")
	podman_podCmd.Flags().StringP("remote", "r", "", "")
	podman_podCmd.Flags().String("root", "", "")
	podman_podCmd.Flags().String("runroot", "", "")
	podman_podCmd.Flags().String("runtime", "", "")
	podman_podCmd.Flags().String("runtime-flag", "", "")
	podman_podCmd.Flags().String("ssh", "", "")
	podman_podCmd.Flags().String("storage-driver", "", "")
	podman_podCmd.Flags().String("storage-opt", "", "")
	podman_podCmd.Flags().Bool("syslog", false, "")
	podman_podCmd.Flags().String("tmpdir", "", "")
	podman_podCmd.Flags().String("transient-store", "", "")
	podman_podCmd.Flags().String("url", "", "")
	podman_podCmd.Flags().String("volumepath", "", "")
	podmanCmd.AddCommand(podman_podCmd)
}
