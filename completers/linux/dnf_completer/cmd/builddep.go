package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var builddepCmd = &cobra.Command{
	Use:     "builddep",
	Short:   "Install build dependencies for package or spec file",
	GroupID: "commands",
	Aliases: []string{"build-dep"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(builddepCmd).Standalone()

	builddepCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	builddepCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	builddepCmd.Flags().StringP("define", "D", "", "Define the RPM macro named \"macro\" to the value \"expr\" when parsing spec files")
	builddepCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	builddepCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	builddepCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	builddepCmd.Flags().String("with", "", "Enable conditional build OPTION when parsing spec files")
	builddepCmd.Flags().String("without", "", "Disable conditional build OPTION when parsing spec files")
	rootCmd.AddCommand(builddepCmd)
}
