package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
