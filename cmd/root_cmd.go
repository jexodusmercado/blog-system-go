package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "blog-system-go",
	Run: func(cmd *cobra.Command, args []string) {
		// migrate(cmd, args)
		serve(cmd.Context())
	},
}

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command {
	rootCmd.AddCommand(&serveCmd)

	return &rootCmd
}