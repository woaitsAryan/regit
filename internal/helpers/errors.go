package helpers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func ThrowError(message string, e error, path string) {
	if e != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true).
			Padding(1)
		message := fmt.Sprintf("%s\n\nPlease report this issue at https://github.com/woaitsAryan/regit/issues so it can be fixed.\n\nError callstack: %s at %s", message, e, path)
		styledMessage := errorStyle.Render(message)

		fmt.Println(styledMessage)
	}
}
