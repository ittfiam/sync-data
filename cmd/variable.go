package cmd

import (
	"fmt"

	"os"

	"sync-mysql/sync"
	"github.com/spf13/cobra"
)

func init() {

}

func NewVariableCmd() *cobra.Command {

	command := &cobra.Command{
		Use:   "variable",
		Short: "manage variables",
		Run: func(cmd *cobra.Command, args []string) {

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Print(
					err.Error())
				return
			}

			sync.Format(variable, os.Stdout, 0, 0)
		},
	}

	command.AddCommand(
		&cobra.Command{
			Use:   "set",
			Short: "set variable (sync-mysql variable set <key> <value>)",
			Run: func(cmd *cobra.Command, args []string) {

				if len(args) != 2 {
					fmt.Println("args must key value.")
				}

				variable, err := sync.NewVariables()

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				variable.Set(args[0], args[1])

				err = variable.Save()

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Printf("%s:%s\n", args[0], args[1])

			},
		},
	)

	return command
}
