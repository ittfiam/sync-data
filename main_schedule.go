package main


import (
	"sync-mysql/cmd"
	"os"
	"os/signal"
	"fmt"
	"github.com/robfig/cron"
)

/**
定时任务，每天(时)增量同步数据
秒、分、时、日、月、星期（非必填）
 */

var (
	root = cmd.NewRootCmd()
)


func main(){
	root.AddCommand(
		cmd.NewSchemaCmd(),
		cmd.NewSyncCmd(),
	)
	c := cron.New()
	// 添加每日创建增量同步描述文件
	c.AddFunc("0 0 4 * * *",schedulerSyncFile)
	// 添加3个点定时执行增量同步任务
	c.AddFunc("0 15 5 * * *", SchedulerIncrement)
	c.AddFunc("0 15 13 * * *", SchedulerIncrement)
	c.AddFunc("0 15 18 * * *", SchedulerIncrement)
	c.Start()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	signalOccur := <-signals
	fmt.Println("Signal occured, signal:", signalOccur.String())
}


func schedulerSyncFile(){
	fmt.Println("Start")
	root.SetArgs([]string{"describe","plan","--source=$source","--target=$target"})
	err := root.Execute()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("end")

}

func SchedulerIncrement(){
	fmt.Println("Start")
	root.SetArgs([]string{"sync","plan","--command=$command","--datax=$datax"})
	err := root.Execute()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("end")
}
