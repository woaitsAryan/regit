package tools

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecuteRewrite(source string, command []string, flags map[string]bool) {
	bashScriptPath := "./git-filter-repo"

    args := append([]string{"--source", source, "--target", source}, command...)

	cmd := exec.Command(bashScriptPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if(flags["quiet"]){
		return
	}
	if(flags["verbose"]){
		fmt.Println(string(output))
	}
	fmt.Println("Done!")
}