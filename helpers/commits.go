package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"github.com/woaitsAryan/regit/models"
)

func GetTotalCommits(flags models.Flags) int {
	totalCommits := fmt.Sprintf("git -C %s log --oneline | wc -l", flags.Source)

	cmd := exec.Command("bash", "-c", totalCommits)
	output, err := cmd.CombinedOutput()
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
