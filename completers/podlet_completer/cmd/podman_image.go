package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var podman_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(podman_imageCmd).Standalone()

	podman_imageCmd.Flags().String("cgroup-manager", "", "")
	podman_imageCmd.Flags().String("config", "", "")
	podman_imageCmd.Flags().String("conmon", "", "")
	podman_imageCmd.Flags().String("connection", "", "")
	podman_imageCmd.Flags().String("events-backend", "", "")
	podman_imageCmd.Flags().String("hooks-dir", "", "")
	podman_imageCmd.Flags().String("identity", "", "")
	podman_imageCmd.Flags().String("imagestore", "", "")
	podman_imageCmd.Flags().String("log-level", "", "")
	podman_imageCmd.Flags().String("module", "", "")
	podman_imageCmd.Flags().String("network-cmd-path", "", "")
	podman_imageCmd.Flags().String("network-config-dir", "", "")
	podman_imageCmd.Flags().String("out", "", "")
	podman_imageCmd.Flags().StringP("remote", "r", "", "")
	podman_imageCmd.Flags().String("root", "", "")
	podman_imageCmd.Flags().String("runroot", "", "")
	podman_imageCmd.Flags().String("runtime", "", "")
	podman_imageCmd.Flags().String("runtime-flag", "", "")
	podman_imageCmd.Flags().String("ssh", "", "")
	podman_imageCmd.Flags().String("storage-driver", "", "")
	podman_imageCmd.Flags().String("storage-opt", "", "")
	podman_imageCmd.Flags().Bool("syslog", false, "")
	podman_imageCmd.Flags().String("tmpdir", "", "")
	podman_imageCmd.Flags().String("transient-store", "", "")
	podman_imageCmd.Flags().String("url", "", "")
	podman_imageCmd.Flags().String("volumepath", "", "")
	podmanCmd.AddCommand(podman_imageCmd)
}
