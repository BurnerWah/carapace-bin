package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionPeers() carapace.Action {
	return carapace.ActionExecCommand("tailscale", "status")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 && fields[0] != "#" {
				vals = append(vals, fields[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
