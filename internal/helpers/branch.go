package helpers

import (
	"os/exec"
	"strings"
)

func branchExists(branchName string, source string) (bool, error) {
    cmd := exec.Command("git", "-C", source, "branch", "--list")
    output, err := cmd.Output()
    if err != nil {
        return false, err
    }

    branches := strings.Split(string(output), "\n")
    for _, branch := range branches {
        cleanBranch := strings.TrimSpace(strings.Replace(branch, "*", "", -1))
        if cleanBranch == branchName {
            return true, nil
        }
    }
    return false, nil
}