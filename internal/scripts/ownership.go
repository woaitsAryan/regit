package scripts

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/woaitsAryan/regit/internal/helpers"
	"github.com/woaitsAryan/regit/internal/models"
)

func Owngit(flags models.Flags ) {
	if(!flags.Quiet){
		fmt.Println("Owning Git..")
	}

	nameCmd := exec.Command("git", "config", "user.name")
    nameOut, err := nameCmd.Output()
    if err != nil {
		helpers.ThrowError("Error fetching Git username", err, "internal/scripts/ownership.go")
        return
    }
	name := strings.TrimSpace(string(nameOut))

    emailCmd := exec.Command("git", "config", "user.email")
    emailOut, err := emailCmd.Output()
    if err != nil {
		helpers.ThrowError("Error fetching Git email", err, "internal/scripts/ownership.go")
        return
    }
	email := strings.TrimSpace(string(emailOut))
	if(flags.Verbose){
		fmt.Printf("Local credentials found! %s <%s>\n", name, email)
	}
	ownCmd := []string{
		"--email-callback",
		fmt.Sprintf("return b\"%s\"", email),
		"--name-callback",
		fmt.Sprintf("return b\"%s\"", name),
		"--force",
	}	
	helpers.ExecuteRewrite(ownCmd, flags)
}

func Blamegit(name string, email string, flags models.Flags){
	if(!flags.Quiet){
		fmt.Printf("Blaming Git to %s <%s>..\n", name, email)
	}
	blameCmd := []string{
		"--email-callback",
		fmt.Sprintf("return b\"%s\"", email),
		"--name-callback",
		fmt.Sprintf("return b\"%s\"", name),
		"--force",
	}		
	helpers.ExecuteRewrite(blameCmd, flags)
}