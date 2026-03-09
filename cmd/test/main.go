package main

import (
	"fmt"
	"os"

	lib "github.com/wagoodman/go-monorepo-test"
)

// these can be set via ldflags at build time
var (
	version   = lib.Version
	buildDate = lib.BuildDate
	commit    = lib.CommitHash
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", buildDate)
		fmt.Printf("Commit: %s\n", commit)
		return
	}

	fmt.Println(lib.Hello())
}
