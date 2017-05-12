package sync

import (

	"path/filepath"
	"fmt"
	"time"
	"strings"
)

type Plan struct {
	tasks    []*Task
	Describe *Describe
	Context  *PlanContext
}

type Task struct {
	Compiled *CompiledJob

	Sync []*Table
	Plan *Plan
}

type PlanMode int

const PlanModeSync PlanMode = 0
const PlanModeUpdate PlanMode = 1

type PlanContext struct {
	Name           string `json:"name"`
	Mode           PlanMode `json:"mode"`
	StartTime  time.Time `json:"start_time"`
	EndTime  time.Time `json:"end_time"`
	Param map[string]string `json:"param"`
	UseTime float64
	SyncTableCount int
}



func (ctx *PlanContext) Save() error {

	return SaveAssetAsJSON(
		filepath.Join("context", ctx.Name),
		ctx)

}

func (ctx *PlanContext) Load() error{
	return ReadAssetAsJSON(
		filepath.Join("context", "context.json"),
		ctx)

}

func (ctx *PlanContext) Start() {
	ctx.StartTime = time.Now()

}


func (ctx *PlanContext) End() {
	ctx.EndTime = time.Now()
	ctx.UseTime = ctx.EndTime.Sub(ctx.StartTime).Seconds()
	if ctx.Param == nil{
		ctx.Param = make(map[string]string,0)
	}
	t := ctx.StartTime.Format("2006-01-02T15:04:05")
	s := strings.Split(t,"T")
	ctx.Param["last_update_date"] = s[0]
	ctx.Param["last_update_time"] = s[1]
	ctx.Save()

}


func (ctx *PlanContext) GetDataxParam() string{

	if ctx.Param == nil{
		return ""
	}
	s := make([]string,0)
	for key,value := range ctx.Param{
		v :=fmt.Sprintf("-D%s=%s",key,value)
		s = append(s,v)
	}
	p := strings.Join(s," ")
	return fmt.Sprintf("-p \"%s\"",p)

}

