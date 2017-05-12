package cmd

import (
	"fmt"

	"os"

	"sync-mysql/sync"
	"github.com/spf13/cobra"
)

func NewSchemaCmd() *cobra.Command {

	var prefix string
	var table, column bool

	skips := make([]string, 0)

	command := &cobra.Command{
		Use:     "schema",
		Short:   "show mysql schema",
		Example: "schema username:password@protocol(address)/dbname\nschema #varialbe",
		Run: func(cmd *cobra.Command, args []string) {

			mysql := args[0]

			if len(args) != 1 {
				fmt.Println("args reqired\nuse: dark-sync schema <sql-str/%%variable>")
				return
			}

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&mysql)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			info := sync.NewSchemaInfo()
			info.Prefix = prefix
			info.AddSkips(skips...)

			schema, err := sync.NewSchemaFromMysql(mysql, info)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			flags := 0

			if table {
				flags |= sync.FORMAT_TABLE
			}

			if column {
				flags |= sync.FORMAT_COLUMN
			}

			flags |= sync.FORMAT_DATABASE

			sync.Format(schema, os.Stdout, 0, flags)
		},
	}

	flags := command.Flags()

	flags.BoolVarP(
		&table,
		"table",
		"t",
		true,
		"show tables.")

	flags.BoolVarP(
		&column,
		"column",
		"c",
		false,
		"show columns.")

	flags.StringSliceVar(
		&skips,
		"skips",
		skips,
		"skip db names.")

	flags.StringVarP(
		&prefix,
		"prefix",
		"p",
		"",
		"filter db prefix.")

	return command
}
