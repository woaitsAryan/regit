package helpers

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/woaitsAryan/regit/internal/models"
)

func ExecuteRewrite(command []string, flags models.Flags) {
	if (flags.Branch != "."){
		exists, err := branchExists(flags.Branch, flags.Source)
		if err != nil {
			ThrowError("Error validating branch name", err, "internal/helpers/execute.go")
		}
		if !exists {
			tempMsg := fmt.Sprintf("Branch %s does not exist", flags.Branch)
			tempError := errors.New("unable to find the specified branch")
			ThrowError(tempMsg, tempError, "internal/helpers/execute.go")
		}
		command = append(command, "--refs", flags.Branch)
	}
	args := append([]string{"filter-repo", "--source", flags.Source, "--target", flags.Source}, command...)

	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		ThrowError("Error executing git command", err, "internal/helpers/execute.go")
	}
	if flags.Quiet {
		return
	}
	if flags.Verbose {
		fmt.Println(string(output))
	}
	fmt.Println("Done!")
}
