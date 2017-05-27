package sync

import (
	"encoding/json"

	"fmt"

	"sync-mysql/errors"
)

type SelectorCond map[string]interface{}
type UpdateCond map[string]interface{}

type Selector struct {
	DB    SelectorCond `json:"db"`
	Table SelectorCond `json:"table"`
}

func NewSelector(db SelectorCond, table SelectorCond) *Selector {
	return &Selector{
		DB:    db,
		Table: table,
	}
}

type UpdateSelector struct {
	Selector
	Cond string `json:"cond"`
}

func (selector *UpdateSelector) IsHaveCond() bool {
	return len(selector.Cond) > 0
}

func NewUpdateSelector(db, table SelectorCond, cond string) *UpdateSelector {

	selector := new(UpdateSelector)
	selector.DB = db
	selector.Table = table
	selector.Cond = cond

	return selector
}

type Job struct {
	Name       string          `json:"name"`
	Enable     bool            `json:"enable"`
	Sync       *Selector       `json:"sync"`
	Update     *UpdateSelector `json:"update"`
	Keys       []string        `json:"keys"`
	DB         string          `json:"db"`
	Collection string          `json:"collection"`
	Work *Work
	Sql []string
	Mode string
	// 写入模式组合 reader2writer
	WRName string
}


func NewJobs(describe []byte) ([]*Job, error) {

	type FileContent struct {
		Jobs []*Job `json:"jobs"`
	}

	c := new(FileContent)

	err := json.Unmarshal(describe, c)

	if err != nil {
		return nil, errors.ToFormatError(
			err,
			"unmarsal jobs describe fail.",
		)
	}

	return c.Jobs, nil

}

func (job *Job) IsSyncDB(name string) bool {

	expr, err := ToExpr(job.Sync.DB, nil)

	if err != nil {
		panic(err.Error())
	}

	return expr.Eval(name)
}

func (job *Job) SaveSql(path string) error{

	if job.Sql == nil || len(job.Sql) <= 0{
		return nil
	}
	bytes := make([][]byte,0)
	for _,s := range job.Sql{
		bytes = append(bytes,[]byte(s))
	}

	sb := BytesCombine("\n",bytes...)

	return SaveAssetFile(path,sb)

}


type CompiledSelector struct {
	DB    Expr
	Table Expr
}

func (selector *CompiledSelector) Matched(db, table string) bool {
	return selector.DB.Eval(db) && selector.Table.Eval(table)
}

type CompiledUpdateSelector struct {
	CompiledSelector
	Cond string
}

type CompiledJob struct {
	Sync   *CompiledSelector
	Update *CompiledUpdateSelector
	Job    *Job
}

func NewCompileJob(index int, job *Job) (complied *CompiledJob, err error) {

	complied = new(CompiledJob)
	complied.Job = job
	complied.Sync = new(CompiledSelector)
	complied.Update = new(CompiledUpdateSelector)
	complied.Update.Cond = job.Update.Cond

	complied.Sync.DB, err = ToExpr(
		job.Sync.DB,
		NewExprStack(
			fmt.Sprintf(
				"%v.%s", index, "job.sync.db")))

	if err != nil {
		err = errors.ToFormatError(
			err,
			"make expr fail.",
		)

		return
	}

	complied.Sync.Table, err = ToExpr(
		job.Sync.Table,
		NewExprStack(
			fmt.Sprintf(
				"%v.%s", index, "job.sync.table")))

	if err != nil {
		err = errors.ToFormatError(
			err,
			"make expr fail.",
		)

		return
	}

	complied.Update.DB, err = ToExpr(
		job.Update.DB,
		NewExprStack(
			fmt.Sprintf(
				"%v.%s", index, "job.update.db")))

	if err != nil {
		err = errors.ToFormatError(
			err,
			"make expr fail.",
		)

		return
	}

	complied.Update.Table, err = ToExpr(
		job.Update.Table,
		NewExprStack(
			fmt.Sprintf(
				"%v.%s", index, "job.update.table")))

	if err != nil {
		err = errors.ToFormatError(
			err,
			"make expr fail.",
		)

		return
	}

	return
}

func Complie(jobs []*Job) (compiled []*CompiledJob, err error) {

	compiled = make([]*CompiledJob, len(jobs))

	for index, job := range jobs {
		compiled[index], err = NewCompileJob(index, job)

		if err != nil {
			return
		}
	}

	return
}
