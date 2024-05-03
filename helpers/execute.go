package helpers

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecuteRewrite(source string, command []string, flags map[string]bool) {
	args := append([]string{"filter-repo", "--source", source, "--target", source}, command...)

	cmd := exec.Command("git", args...)
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
