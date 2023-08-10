package main

import (
	"fmt"

	"github.com/spfl13/cobra"
)

func initVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show the tool version information",
		Run:   version,
	}
	return versionCmd
}
func version(cmd *cobra.Command, args []string) {
	fmt.Println("v1.0.0")
}
