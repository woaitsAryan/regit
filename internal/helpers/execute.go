package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"github.com/woaitsAryan/regit/internal/models"

)

func ExecuteRewrite(command []string, flags models.Flags) {
	args := append([]string{"filter-repo", "--source", flags.Source, "--target", flags.Source}, command...)

	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if flags.Quiet {
		return
	}
	if flags.Verbose {
		fmt.Println(string(output))
	}
	fmt.Println("Done!")
}
