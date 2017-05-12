package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {

	command := &cobra.Command{
		Use: "dark-sync",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		}}

	return command
}
