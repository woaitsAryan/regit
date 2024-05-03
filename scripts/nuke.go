package scripts

import (
	"fmt"
	"github.com/woaitsAryan/regit/helpers"
)

func NukeGit(path string, file string, flags map[string]bool) {

	fmt.Println("Now I am become regit, destroyer of Git histories")

	nukeCmd := []string{
		"--invert-paths",
		"--path",
		file,
		"--force",
	}

	helpers.ExecuteRewrite(path, nukeCmd, flags)
}