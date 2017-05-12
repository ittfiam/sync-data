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
	"errors"
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


	f := filepath.Join("describes","sql",fileName)
	ok, err := sync.AssetExists(f)

	if err != nil {
		return err
	}

	if !ok{
		fmt.Println("文件不存在",fileName)
		return errors.New(fmt.Sprintf("文件不存在:%s" , fileName))
	}


	s,err:=sync.ReadAssetAsString(f)
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

	var command,datax string
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

			err = variable.GetValue(&command,&datax)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fileList,_,err := sync.ReadFileList(filepath.Join("describes","init"))

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
				subFileList,parent,err:=sync.ReadFileList(filepath.Join("describes","init",f.Name()))
				if err != nil{
					fmt.Println(err.Error())
					return
				}
				for _,subF := range subFileList{

					if subF.IsDir(){
						continue
					}

					cp :=exec.Command(command,datax,filepath.Join(parent,subF.Name()))
					run(cp)
				}


			}

		},
	}

	flags := c.Flags()

	flags.StringVar(
		&command,
		"command",
		"",
		"use datax sync data command (value or $command)")

	flags.StringVar(
		&datax,
		"datax",
		"",
		"use datax sync data command (value or $datax)")


	return c
}


func syncPlanCmd() *cobra.Command {

	var command,datax string
	var notifyUrl = "$notifyUrl"
	c := &cobra.Command{
		Use:     "plan",
		Short:   "execute sync plan",
		Example: "sync-mysql sync plan",
		Run: func(cmd *cobra.Command, args []string) {

			//	执行同步命令（init）

			variable, err := sync.NewVariables()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = variable.GetValue(&command,&datax,&notify)

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
				subFileList,parent,err:=sync.ReadFileList(filepath.Join("describes","init",f.Name()))
				if err != nil{
					fmt.Println(err.Error())
					return
				}
				for _,subF := range subFileList{

					if subF.IsDir(){
						continue
					}

					cp :=exec.Command(command,datax,pc.GetDataxParam(),filepath.Join(parent,subF.Name()))
					run(cp)
					pc.SyncTableCount = pc.SyncTableCount + 1
				}


			}
		//	save context
			pc.End()
			go notify.Notify(notifyUrl)

		},
	}

	flags := c.Flags()

	flags.StringVar(
		&command,
		"command",
		"",
		"use datax sync data command (value or $command)")

	flags.StringVar(
		&datax,
		"datax",
		"",
		"use datax sync data command (value or $datax)")


	return c
}

func syncCreateCmd() *cobra.Command {

	var file,target string
	c := &cobra.Command{
		Use:     "runsql",
		Short:   "drop and create table",
		Example: "sync-mysql create table",
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
			err = v.GetValue(&file,&target)
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
				fileList,_,err := sync.ReadFileList(filepath.Join("describes","sql"))

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
					err := runPreSql(cn,f.Name())
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
		"",
		"use create table sql file(all,为所有)")

	flags.StringVar(
		&target,
		"target",
		"",
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
