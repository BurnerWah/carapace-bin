package dnf

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type repo struct {
	Id        string
	Name      string
	IsEnabled bool `json:"is_enabled"`
}

// ActionRepos completes dnf repos
func ActionRepos() carapace.Action {
	return carapace.ActionExecCommand("dnf", "-C", "repo", "list", "--all", "--json")(func(output []byte) carapace.Action {
		var repos []repo
		if err := json.Unmarshal(output, &repos); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0, len(repos)*2)
		for _, repo := range repos {
			vals = append(vals, repo.Id, repo.Name)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
