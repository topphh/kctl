package cmd

import (
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Show resource usage metrics",
	Long: `The 'top' command shows CPU and memory usage for Kubernetes resources.

Usage examples:

  kctl top service        # Show CPU and memory usage by service

This is useful for quickly identifying performance bottlenecks
or monitoring resource consumption over time.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
