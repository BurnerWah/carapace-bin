package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(composeCmd).Standalone()

	composeCmd.Flags().Bool("kube", false, "Create a Kubernetes YAML file for a pod instead of separate containers")
	composeCmd.Flags().Bool("pod", false, "Create a `.pod` file and link it with each `.container` file")
	rootCmd.AddCommand(composeCmd)
}
