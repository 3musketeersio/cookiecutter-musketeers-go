package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version    string
	buildDate  string
	commitHash string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the {{cookiecutter.git_repo_name}} version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf(`{{ cookiecutter.git_repo_name }}:
version     : %s
build date  : %s
git hash    : %s
go version  : %s
go compiler : %s
platform    : %s/%s
`, version, buildDate, commitHash, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
		return nil
	},
}

func init() {
	// init is also used to add any flag and add the command to the rootCmd
	rootCmd.AddCommand(versionCmd)
}
