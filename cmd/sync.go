package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os/exec"
	"sync-mysql/sync"
	"path/filepath"
	"bufio"
	"io"
	"strings"
	"sync-mysql/notify"
)


func run(cmd *exec.Cmd) bool{

	fmt.Println(cmd.Args)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}


func runPreSql(conn string, fileName string) error{

	s,err:=sync.ReadAsString(fileName)
	if err != nil{
		fmt.Println(err.Error())
		return err
	}
	sr := strings.Split(s,";\n")
	err = sync.ExecSql(conn,sr...)
	if err != nil{

		fmt.Println(err)
		return err
	}
	return nil
}

func syncInitCmd() *cobra.Command {

	var command sync.SyncCmdParam
	c := &cobra.Command{
		Use:     "init",
		Short:   "execute sync",
		Example: "sync-mysql sync init",
		Run: func(cmd *cobra.Command, args []string) {

		//	执行同步命令（init）

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&command.Command,&command.Mode,&command.DataxPath,&command.EnableNotify,&command.NotifyUrl)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fileList,_,err := sync.ReadFileList(filepath.Join("describes",command.Mode,"init"))

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if len(fileList) <= 0 {
				return
			}


			for _,f :=range fileList{

				if !f.IsDir(){
					continue
				}
				fp := filepath.Join("describes",command.Mode,"init",f.Name())
				subFileList,parent,err:=sync.ReadFileList(fp)
				if err != nil{
					fmt.Println(err.Error())
					return
				}
				for _,subF := range subFileList{

					if subF.IsDir(){
						continue
					}

					cp :=exec.Command(command.Command,command.DataxPath,filepath.Join(parent,subF.Name()))
					run(cp)
				}


			}

		},
	}

	flags := c.Flags()

	flags.StringVar(
		&command.Command,
		"command",
		"$command",
		"use datax sync data command (value or $command)")

	flags.StringVar(
		&command.DataxPath,
		"dataxpath",
		"$dataxPath",
		"use datax sync data command (value or $dataxPath)")

	flags.StringVar(
		&command.Mode,
		"mode",
		"$mode",
		"use datax sync data command (value or $mode)")


	return c
}


func syncPlanCmd() *cobra.Command {

	var command sync.SyncCmdParam
	c := &cobra.Command{
		Use:     "plan",
		Short:   "execute sync plan",
		Example: "sync-mysqlplugin sync plan",
		Run: func(cmd *cobra.Command, args []string) {

			//	执行同步命令（init）

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&command.Command,&command.Mode,&command.DataxPath,&command.EnableNotify,&command.NotifyUrl)


			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fileList,_,err := sync.ReadFileList(filepath.Join("describes","increment"))

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if len(fileList) <= 0 {
				return
			}
			// 读取运行变量
			var pc sync.PlanContext

			err =pc.Load()
			if err != nil{
				fmt.Println(err.Error())
				return
			}

			pc.Start()

			for _,f :=range fileList{

				if !f.IsDir(){
					continue
				}
				subFileList,parent,err:=sync.ReadFileList(filepath.Join("describes",command.Mode,"plan",f.Name()))
				if err != nil{
					fmt.Println(err.Error())
					return
				}
				for _,subF := range subFileList{

					if subF.IsDir(){
						continue
					}

					cp :=exec.Command(command.Command,command.DataxPath,pc.GetDataxParam(),filepath.Join(parent,subF.Name()))
					run(cp)
					pc.SyncTableCount = pc.SyncTableCount + 1
				}


			}
			//save context
			pc.End()
			if command.EnableNotify == "true"{
				go notify.NotifyBoss(command.NotifyUrl)
			}

		},
	}

	flags := c.Flags()

	flags.StringVar(
		&command.Command,
		"command",
		"$command",
		"use datax sync data command (value or $command)")

	flags.StringVar(
		&command.DataxPath,
		"dataxpath",
		"$dataxPath",
		"use datax sync data command (value or $dataxPath)")

	flags.StringVar(
		&command.Mode,
		"mode",
		"$mode",
		"use datax sync data command (value or $mode)")

	flags.StringVar(
		&command.EnableNotify,
		"enableNotify",
		"$enableNotify",
		"use datax sync data command (value or $enableNotify)")


	flags.StringVar(
		&command.NotifyUrl,
		"notifyUrl",
		"$notifyUrl",
		"use datax sync data command (value or $notifyUrl)")

	return c
}

func syncCreateCmd() *cobra.Command {

	var file,target,mode string
	c := &cobra.Command{
		Use:     "runsql",
		Short:   "drop and create table",
		Example: "sync-mysqlplugin create table",
		Run: func(cmd *cobra.Command, args []string) {

			// 执行 初始化表命令

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&file)

			if err != nil {
				fmt.Println(err.Error())
				return
			}


			// 获取连接信息
			v,err := sync.NewVariables()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = v.GetValue(&file,&target,&mode)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			conninfo,err := sync.ParseScheme(target)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			cn := conninfo.ToGoMysqlAndDB()

			if file == "all"{

				fp := filepath.Join("describes",mode,"sql")
				fileList,parent,err := sync.ReadFileList(fp)

				if err != nil {
					fmt.Println(err.Error())
					return
				}
				if len(fileList) <= 0 {
					return
				}

				for _,f :=range fileList{

					if f.IsDir(){
						continue
					}
					err := runPreSql(cn,filepath.Join(parent,f.Name()))
					if err != nil{
						fmt.Println(err)
						return
					}
				}
			}else {
				err := runPreSql(cn,file)
				if err != nil{
					fmt.Println(err)
					return
				}
			}



		},
	}

	flags := c.Flags()

	flags.StringVar(
		&file,
		"file",
		"all",
		"use create table sql file(all,为所有)")

	flags.StringVar(
		&target,
		"target",
		"$target",
		"use sync db scheme")

	flags.StringVar(
		&mode,
		"mode",
		"$mode",
		"use sync db scheme")


	return c
}

func NewSyncCmd() *cobra.Command {

	command := &cobra.Command{
		Use:   "sync",
		Short: "manage sync",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	command.AddCommand(
		syncInitCmd(),
		syncPlanCmd(),
		syncCreateCmd(),
	)

	return command

}
