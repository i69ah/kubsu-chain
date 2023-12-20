package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Major = "1"
const Minor = "0"
const Fix = "0"
const Verbal = "TBB-net"

// Configured via -ldflags during build
var GitCommit string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Describes version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Version: %s.%s.%s-alpha %s %s", Major, Minor, Fix, shortGitCommit(GitCommit), Verbal))
	},
}

func shortGitCommit(fullGitCommit string) string {
	shortCommit := ""
	if len(fullGitCommit) >= 6 {
		shortCommit = fullGitCommit[0:6]
	}

	return shortCommit
}
