package main

import (
	"github.com/spfl13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "exp [Command]",
	}
	rootCmd.AddCommand(initVersionCmd())
	rootCmd.AddCommand(initRunCmd())
	rootCmd.Execute()
}
