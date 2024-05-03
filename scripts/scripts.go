package scripts

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/woaitsAryan/regit/helpers"
)

func Owngit(path string, flags map[string]bool ) {
	if(!flags["quiet"]){
		fmt.Println("Owning Git..")
	}

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
	if(flags["verbose"]){
		fmt.Printf("Local credentials found! %s <%s>\n", name, email)
	}
	ownCmd := []string{
		"--email-callback",
		fmt.Sprintf("return b\"%s\"", email),
		"--name-callback",
		fmt.Sprintf("return b\"%s\"", name),
		"--force",
	}	
	helpers.ExecuteRewrite(path, ownCmd, flags)
}

func Blamegit(path string, name string, email string, flags map[string]bool ){
	if(!flags["quiet"]){
		fmt.Printf("Blaming Git to %s <%s>..\n", name, email)
	}
	blameCmd := []string{
		"--email-callback",
		fmt.Sprintf("return b\"%s\"", email),
		"--name-callback",
		fmt.Sprintf("return b\"%s\"", name),
		"--force",
	}		
	helpers.ExecuteRewrite(path, blameCmd, flags)
}