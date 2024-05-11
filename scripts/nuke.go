package scripts

import (
	"fmt"
	"github.com/woaitsAryan/regit/helpers"
	"github.com/woaitsAryan/regit/models"
)

func NukeGit(file string, flags models.Flags) {

	fmt.Println("Now I am become regit, destroyer of Git histories")

	nukeCmd := []string{
		"--invert-paths",
		"--path",
		file,
		"--force",
	}

	helpers.ExecuteRewrite(nukeCmd, flags)
}