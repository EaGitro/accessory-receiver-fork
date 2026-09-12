package main

import (
	"os"

	"github.com/spf13/afero"

	"github.com/EaGitro/accessory-receiver-fork/cmd"
)

func main() {
	cmd.Execute(afero.NewOsFs(), os.Args)
}
