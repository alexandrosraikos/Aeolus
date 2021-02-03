package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "anchosctl",
		Short: "A blazing fast cloud infrastructure semantic orchestrator.",
		Long: `Anchor is a semantic orchestrator for cloud
		infrastructure applications.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
