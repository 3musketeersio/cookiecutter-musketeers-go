package main

import (
	"{{cookiecutter.app_go_package_path}}/cmd"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

func main() {
	cmd.Execute(version, commitHash, buildDate)
}
