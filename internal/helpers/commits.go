package helpers

import (
	"fmt"
	"github.com/woaitsAryan/regit/internal/models"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetTotalCommits(flags models.Flags) int {
	cmd := []string{"git", "-C", flags.Source, "rev-list", "--count"}

	if flags.Branch != "." {
		exists, err := branchExists(flags.Branch, flags.Source)
		if err != nil {
			log.Fatalf("branchExists() failed with %s\n", err)
		}
		if !exists {
			log.Fatalf("Branch %s does not exist\n", flags.Branch)
		}
		cmd = append(cmd, flags.Branch)
	} else {
		cmd = append(cmd, "HEAD")
	}

	execCmd := exec.Command(cmd[0], cmd[1:]...)
	output, err := execCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outputStr := string(output)
	outputInt, err := strconv.Atoi(strings.TrimSpace(outputStr))
	if err != nil {
		log.Fatalf("strconv.Atoi() failed with %s\n", err)
	}
	if flags.Quiet {
		return outputInt
	}
	fmt.Printf("Total number of commits: %d\n", outputInt)

	return outputInt
}
