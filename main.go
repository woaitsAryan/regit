package main

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/initializers"
	"github.com/woaitsAryan/regit/setup"
)

var rootCmd = &cobra.Command{Use: "regit"}

func init(){
	initializers.LoadEnv()
	initializers.CheckCommand()
}

func main() {
	rootCmd.AddCommand(setup.OwnGitCommand)
	rootCmd.AddCommand(setup.BlameGitCommand)
	rootCmd.AddCommand(setup.BlameLinusCommand)
	rootCmd.AddCommand(setup.RetimeCommand)
	rootCmd.AddCommand(setup.NukeGitCommand)
	cobra.CheckErr(rootCmd.Execute())
}