package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var declarationFileRootFlag string
var tagsRootFlag []string

var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.git_repo_name}}",
	Short: "{{cookiecutter.git_repo_name}} does something wonderful",
	Long: `{{cookiecutter.git_repo_name}} does something wonderful and even more!
Usage example
  print out version information
    $ {{cookiecutter.git_repo_name}} version`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute(version, commitHash, buildDate string) {
	version = version
	commitHash = commitHash
	buildDate = buildDate
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
