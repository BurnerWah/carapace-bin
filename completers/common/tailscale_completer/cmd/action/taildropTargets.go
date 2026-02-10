package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionTaildropTargets() carapace.Action {
	return carapace.ActionExecCommand("tailscale", "file", "cp", "--targets")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[1])
			}
		}

		return carapace.ActionValues(vals...)
	})
}
