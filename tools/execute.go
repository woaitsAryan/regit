package tools

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecuteRewrite(source string, command []string) {
	bashScriptPath := "./git-filter-repo"

    args := append([]string{"--source", source, "--target", source}, command...)

	cmd := exec.Command(bashScriptPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(output))
}