package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podman_kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podman_kubeCmd).Standalone()

	podman_kubeCmd.Flags().String("cgroup-manager", "", "")
	podman_kubeCmd.Flags().String("config", "", "")
	podman_kubeCmd.Flags().String("conmon", "", "")
	podman_kubeCmd.Flags().String("connection", "", "")
	podman_kubeCmd.Flags().String("events-backend", "", "")
	podman_kubeCmd.Flags().String("hooks-dir", "", "")
	podman_kubeCmd.Flags().String("identity", "", "")
	podman_kubeCmd.Flags().String("imagestore", "", "")
	podman_kubeCmd.Flags().String("log-level", "", "")
	podman_kubeCmd.Flags().String("module", "", "")
	podman_kubeCmd.Flags().String("network-cmd-path", "", "")
	podman_kubeCmd.Flags().String("network-config-dir", "", "")
	podman_kubeCmd.Flags().String("out", "", "")
	podman_kubeCmd.Flags().StringP("remote", "r", "", "")
	podman_kubeCmd.Flags().String("root", "", "")
	podman_kubeCmd.Flags().String("runroot", "", "")
	podman_kubeCmd.Flags().String("runtime", "", "")
	podman_kubeCmd.Flags().String("runtime-flag", "", "")
	podman_kubeCmd.Flags().String("ssh", "", "")
	podman_kubeCmd.Flags().String("storage-driver", "", "")
	podman_kubeCmd.Flags().String("storage-opt", "", "")
	podman_kubeCmd.Flags().Bool("syslog", false, "")
	podman_kubeCmd.Flags().String("tmpdir", "", "")
	podman_kubeCmd.Flags().String("transient-store", "", "")
	podman_kubeCmd.Flags().String("url", "", "")
	podman_kubeCmd.Flags().String("volumepath", "", "")
	podmanCmd.AddCommand(podman_kubeCmd)
}
