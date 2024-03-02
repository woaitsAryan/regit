package tools

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func ExecuteRewrite(source string, command []string, flags map[string]bool) {
	bashScriptPath := "./git-filter-repo"

	args := append([]string{"--source", source, "--target", source}, command...)

	cmd := exec.Command(bashScriptPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if flags["quiet"] {
		return
	}
	if flags["verbose"] {
		fmt.Println(string(output))
	}
	fmt.Println("Done!")
}

func getTotalCommits(source string, flags map[string]bool) int {
	totalCommits := fmt.Sprintf("git -C %s log --oneline | wc -l", source)

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
	if flags["quiet"] {
		return outputInt
	}
	fmt.Printf("Total number of commits: %d\n", outputInt)

	return outputInt
}