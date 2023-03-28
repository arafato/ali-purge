package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var verbose bool

	command := &cobra.Command{
		Use:   "ali-purge",
		Short: "ali-purge removes every resource from your Alibaba Cloud account.",
		Long:  `A tool which removes every resource from an Alibaba Cloud account.  Use it with caution, since it cannot distinguish between production and non-production.`,
	}

	command.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false,
		"Enables debug output.")

	command.AddCommand(NewVersionCommand())
	return command
}
