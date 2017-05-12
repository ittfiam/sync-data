package sync

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	FORMAT_DATABASE = 1
	FORMAT_TABLE    = 1 << 1
	FORMAT_COLUMN   = 1 << 2

	FORMAT_SQLS = 1 << 3
)

func formatFlags(flags int) int {

	if flags == 0 {
		flags = FORMAT_DATABASE | FORMAT_TABLE | FORMAT_COLUMN
	}

	if haveFlag(flags, FORMAT_TABLE) {
		flags = flags | FORMAT_DATABASE
	}

	if haveFlag(flags, FORMAT_COLUMN) {
		flags = flags | FORMAT_DATABASE | FORMAT_TABLE
	}

	return flags
}

func haveFlag(flags, flag int) bool {
	return flags&flag == flag
}

func Format(n interface{}, w io.Writer, level, flags int) {

	flags = formatFlags(flags)

	space := strings.Repeat(" ", level*4)

	switch node := n.(type) {
	case *Schema:

		if haveFlag(flags, FORMAT_DATABASE) {
			for _, database := range node.GetDBList() {
				Format(database, w, level, flags)
			}
		}

		var db, table int

		for _, database := range node.DataBases {
			db += 1
			table += len(database.Tables)
		}

		fmt.Fprintf(w, "db:%v,table:%v\n", db, table)
	case *DB:
		fmt.Fprintf(w, "%s%s\n", space, node.Name)
		if haveFlag(flags, FORMAT_TABLE) {
			for _, table := range node.GetTableList() {
				Format(table, w, level+1, flags)
			}
		}
	case *Table:
		fmt.Fprintf(w, "%s%s\n", space, node.Name)

		if haveFlag(flags, FORMAT_COLUMN) {
			for _, column := range node.Columns {
				Format(column, w, level+1, flags)
			}
		}
	case *Column:
		fmt.Fprintf(w, "%s%s\n", space, node.Name)
		fmt.Fprintf(w, "%s%sT:%s\n", space, space, node.Type)

		if len(node.Key) > 0 {
			fmt.Fprintf(w, "%s%sK:%s\n", space, space, node.Key)
		}

		if len(node.Comment) > 0 {
			fmt.Fprintf(w, "%s%sC:%s\n", space, space, node.Comment)
		}
	case *DescribeErrors:
		fmt.Fprintf(w, "%sDescribe errors\n", space)
		if node.Error != nil {
			fmt.Fprintf(w, "%Error: %s\n", space, node.Error.Error())
		}

		if len(node.MoreSync) > 0 {

			fmt.Fprintf(w, "%sTable match more sync block.\n", space)

			for _, block := range node.MoreSync {
				Format(block, w, level+1, flags)
			}

		}

		if len(node.MoreUpdate) > 0 {

			fmt.Fprintf(w, "%sTable match more update block.\n", space)

			for _, block := range node.MoreUpdate {
				Format(block, w, level+1, flags)
			}

		}

		if len(node.NoCond) > 0 {
			fmt.Fprintf(w, "%sJob update.cond is empty.\n", space)

			for _, block := range node.NoCond {
				Format(block, w, level+1, flags)
			}
		}

		if len(node.NotMatch) > 0 {

			fmt.Fprintf(w, "%sTable not match any jobs.\n", space)

			for _, block := range node.NotMatch {
				Format(block, w, level+1, flags)
			}
		}

		if len(node.NotMatchJobs) > 0 {
			fmt.Fprintf(w, "%sJob not match any tabke.\n", space)

			for _, block := range node.NotMatchJobs {
				Format(block, w, level+1, flags)
			}
		}

		fmt.Fprintf(w, "%sSync        :%v\n", space, len(node.MoreSync))
		fmt.Fprintf(w, "%sUpdate      :%v\n", space, len(node.MoreUpdate))
		fmt.Fprintf(w, "%sNoCond      :%v\n", space, len(node.NoCond))
		fmt.Fprintf(w, "%sNotMatch    :%v\n", space, len(node.NotMatch))
		fmt.Fprintf(w, "%sNotMatchJobs:%v\n", space, len(node.NotMatchJobs))
	case *errMoreJobs:
		fmt.Fprintf(w, "%sTable:%s.%s\n", space, node.DB, node.Table)

		for _, job := range node.Jobs {
			fmt.Fprintf(w, "%s%sJob:%s\n", space, space, job.Name)
		}

		fmt.Fprintln(w, "")
	case *errNotMatch:
		fmt.Fprintf(w, "%sTable:%s.%s\n", space, node.DB, node.Table)
	case *errNotMatchJob:
		fmt.Fprintf(w, "%sJob:%s\n", space, node.Job.Name)
	case *errNoCond:
		fmt.Fprintf(w, "%sTable:%s.%s\n", space, node.DB, node.Table)
		fmt.Fprintf(w, "%s%sJob:%s\n", space, space, node.Job.Name)
	case *Variables:

		for name, value := range node.Vars {
			fmt.Fprintf(w, "%s%s\t:%s\n", space, name, value)
		}

		fmt.Fprintf(w, "%sVariables total:%v\n", space, len(node.Vars))
	case *Plan:
		for _, task := range node.tasks {
			Format(task, w, level+1, flags)
		}
	case *Task:
		fmt.Fprintf(w, "%sTask:%s\n", space, node.Compiled.Job.Name)
		fmt.Fprintf(w, "%s%sTables:%v\n", space, space, len(node.Sync))

		if haveFlag(flags, FORMAT_SQLS) {
			fmt.Fprintf(w, "%s%sSQLS:\n", space, space)

			for _, table := range node.Sync {
				fmt.Fprintf(w, "%s%s%s%s\n", space, space, space, table.GetSQL())
			}
		}

	case *CompiledJob:

	}
}

func Print(node interface{}) {
	Format(node, os.Stdout, 0, FORMAT_DATABASE|FORMAT_TABLE|FORMAT_COLUMN)
}
