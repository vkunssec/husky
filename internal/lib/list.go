package lib

import (
	"fmt"

	"github.com/vkunssec/husky/internal/tools"
)

func List() {
	tools.LogUnformatted(" List of hooks implemented in the repository:\n\n")
	hooks := tools.ValidHooksWithDescription

	output := ""
	for _, hook := range hooks {
		output += fmt.Sprintf("  - %s\n", hook)
	}
	output += "\n"
	output += "For more information visit: https://github.com/vkunssec/husky\n"
	output += "If you want to add a new hook, please submit a PR.\n"

	tools.LogUnformatted("%s\n", output)
}
