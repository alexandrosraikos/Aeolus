package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "ploryctl",
		Short: "Your computing node's best friend.",
		Long:  `Plory is an OS daemon which keeps your node connected to Anchos. Tip: A pack of plories is called a Stolos.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
