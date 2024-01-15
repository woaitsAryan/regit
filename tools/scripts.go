package tools

import (
	"fmt"
	"os/exec"
	"strings"
)


func Owngit(path string) {
	fmt.Println("owning git..")

	nameCmd := exec.Command("git", "config", "user.name")
    nameOut, err := nameCmd.Output()
    if err != nil {
        fmt.Println("Error getting Git name:", err)
        return
    }
	name := strings.TrimSpace(string(nameOut))

    emailCmd := exec.Command("git", "config", "user.email")
    emailOut, err := emailCmd.Output()
    if err != nil {
        fmt.Println("Error getting Git email:", err)
        return
    }
	email := strings.TrimSpace(string(emailOut))
	ownCmd := []string{
		"--email-callback",
		fmt.Sprintf("return b\"%s\"", email),
		"--name-callback",
		fmt.Sprintf("return b\"%s\"", name),
		"--force",
	}	
	ExecuteRewrite(path, ownCmd)
}

func Blamegit(path string, name string, email string){
	fmt.Println("blaming git..")
	blameCmd := []string{
		"--email-callback",
		fmt.Sprintf("return b\"%s\"", email),
		"--name-callback",
		fmt.Sprintf("return b\"%s\"", name),
		"--force",
	}		
	ExecuteRewrite(path, blameCmd)
}