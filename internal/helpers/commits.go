package helpers

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/woaitsAryan/regit/internal/models"
)

func GetTotalCommits(flags models.Flags) int {
	cmd := []string{"git", "-C", flags.Source, "rev-list", "--count"}

	if flags.Branch != "." {
		exists, err := branchExists(flags.Branch, flags.Source)
		if err != nil {
			ThrowError(fmt.Sprintf("Error fetching branch %s to count commits", flags.Branch), err, "internal/helpers/commits.go")
		}
		if !exists {
			ThrowError(fmt.Sprintf("Couldn't find the branch %s to count commits", flags.Branch), errors.New("unable to find the branch to fetch commits from"), "internal/helpers/commits.go")
		}
		cmd = append(cmd, flags.Branch)
	} else {
		cmd = append(cmd, "HEAD")
	}

	execCmd := exec.Command(cmd[0], cmd[1:]...)
	output, err := execCmd.CombinedOutput()
	if err != nil {
		ThrowError("Error fetching total number of commits", err, "internal/helpers/commits.go")
	}
	outputStr := string(output)
	outputInt, err := strconv.Atoi(strings.TrimSpace(outputStr))
	if err != nil {
		ThrowError("Error converting number of commits to int", err, "internal/helpers/commit.go")
	}
	if flags.Quiet {
		return outputInt
	}
	fmt.Printf("Total number of commits: %d\n", outputInt)

	return outputInt
}
