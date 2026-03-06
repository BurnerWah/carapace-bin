package dnf

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackages completes dnf packages
func ActionPackages(available bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"-C", "repoquery", fmt.Sprintf("%s*", c.Value), "--qf", "%{name}\n"}
		if available {
			args = append(args, "--available")
		} else {
			args = append(args, "--installed")
		}

		return carapace.ActionExecCommand("dnf", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines...)
		})
	})
}
