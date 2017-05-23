package cmd

import (
	"fmt"
	"path/filepath"
	"sync-mysql/sync"
	"github.com/spf13/cobra"
)



func initCreate(param *sync.CommandParam,info *sync.SchemaInfo){

	source,err := param.GetSourceSchema()

	if err != nil{
		return
	}

	schema, err := sync.NewSchemaFromMysql(source.ToGoMysql(), info)

	if err != nil {
		fmt.Println(err.Error())
		return
	}


	describe,err2 := sync.NewDescribeFromSchema(
		param,
		schema,
	)

	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	if describe == nil{
		fmt.Println("生成描述文件错误")
		return
	}

	for _,job := range describe.Jobs{
		err = sync.SaveAssetAsJSON(
			filepath.Join("describes",job.WRName,"init",job.DB, job.Name + ".json"),
			job.Work,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		job.SaveSql(filepath.Join("describes",job.WRName,"sql",job.Name + ".sql"))
	}
}

func incrementCreate(source *sync.ConnectScheme,target *sync.ConnectScheme,info *sync.SchemaInfo){

	schema, err := sync.NewSchemaFromMysql(source.ToGoMysql(), info)

	if err != nil {
		fmt.Println(err.Error())
	}

	describe,err2 := sync.IncrementDescribeFromSchema(
		source,
		target,
		schema,
		)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	for _,job := range describe.Jobs{
		err = sync.SaveAssetAsJSON(
			filepath.Join("describes","increment",job.DB, job.Name + ".json"),
			job.Work,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}


func describeCreate() *cobra.Command {

	param := new(sync.CommandParam)
	skips := make([]string, 0)

	command := &cobra.Command{
		Use:   "init",
		Short: "manage describes",
		Run: func(cmd *cobra.Command, args []string) {

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&param.Source, &param.Target,&param.Reader,&param.Writer)

			if err != nil {
				fmt.Println(err.Error())
				return
			}


			info := sync.NewSchemaInfo()
			info.Prefix = param.Prefix
			info.AddSkips(skips...)

			initCreate(param,info)

		},
	}

	flags := command.Flags()

	flags.StringVar(
		&param.Source,
		"source",
		"$source",
		"use source schema to generate describe (value or $variable)")

	flags.StringVar(
		&param.Target,
		"target",
		"$target",
		"target db to sync data (value or $variable)")

	flags.StringVar(
		&param.Reader,
		"reader",
		"$reader",
		"reader mode  (value or $variable,like mysqlreader)")

	flags.StringVar(
		&param.Writer,
		"writer",
		"$writer",
		"reader mode  (value or $variable,like mysqlwriter)")

	flags.StringVar(
		&param.Prefix,
		"prefix",
		"",
		"schema database name prefix",
	)

	flags.StringSliceVar(
		&skips,
		"skips",
		skips,
		"schema skip database names",
	)


	return command
}

func describePlan() *cobra.Command {

	var source, target, prefix string
	skips := make([]string, 0)


	command := &cobra.Command{
		Use:   "plan",
		Short: "describe to plan",
		Run: func(cmd *cobra.Command, args []string) {

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&source, &target)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			sourceScheme,err1 := sync.ParseScheme(source)
			if err1 != nil {
				fmt.Println(err.Error())
				return
			}

			targetScheme,err2 := sync.ParseScheme(target)
			if err2 != nil {
				fmt.Println(err.Error())
				return
			}

			info := sync.NewSchemaInfo()
			info.Prefix = prefix
			info.AddSkips(skips...)

			incrementCreate(sourceScheme,targetScheme,info)

		},
	}

	flags := command.Flags()

	flags.StringVar(
		&source,
		"source",
		"$source",
		"use mysql schema to generate describe (value or $variable)")

	flags.StringVar(
		&target,
		"target",
		"$target",
		"user db to sync (value or $variable)")

	flags.StringVar(
		&prefix,
		"prefix",
		"",
		"schema database name prefix",
	)

	flags.StringSliceVar(
		&skips,
		"skips",
		skips,
		"schema skip database names",
	)


	return command
}

func NewDescribeCmd() *cobra.Command {

	command := &cobra.Command{
		Use:   "describe",
		Short: "manage describes",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	command.AddCommand(
		describeCreate(),
		describePlan(),
	)

	return command

}
