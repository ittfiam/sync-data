package main

import (
	"sync-mysql/cmd"
)

func main() {

	root := cmd.NewRootCmd()

	root.AddCommand(
		cmd.NewSchemaCmd(),
		cmd.NewSyncCmd(),
		cmd.NewVariableCmd(),
		cmd.NewDescribeCmd(),
	)

	root.Execute()
}
