package sync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync-mysql/errors"
	"strings"
)

type Describe struct {
	Name  string
	MySQL string
	Mongo string
	Jobs  []*Job `json:"jobs"`
}

func NewDescribeFromAsset(relative string) (describe *Describe, err error) {

	describe = new(Describe)

	err = ReadAssetAsJSON(relative, describe)

	return
}




func NewDescribeFromSchema(
	param *CommandParam,
	schema *Schema,
	) (*Describe,error) {


	sourceScheme,err := param.GetSourceSchema()
	if err != nil{

		return nil,err
	}

	describe := &Describe{
		MySQL: sourceScheme.ToGoMysql(),
		Jobs:  make([]*Job, 0),
	}

	ruleC,err := NewRuleConfigs()


	if err != nil{
		return nil,err
	}

	if ruleC == nil{
		fmt.Println("解析rule规则出错")
		return nil,nil
	}


	for _,db := range schema.GetDBList(){
		r := ruleC.GetRule(db.Name)
		if r == nil{
			continue
		}

		tbs := db.GetTableList()

		temp := make(map[*RuleSub]*DataXContext,0)


		for _,tb := range tbs{

			rSub :=r.GetRuleSub(tb.Name)

			if rSub == nil{
				continue
			}

			sr,ok := temp[rSub]
			if !ok {
				sr = NewDataxContext()
				temp[rSub] = sr
				sr.Table = tb
			}
			sr.SourceTable = append(sr.SourceTable,tb.Name)
			// 重命名表
			sr.Sql = strings.Replace(tb.Sql,tb.Name,rSub.TargetTB,1)
			if rSub.Columns == nil || len(rSub.Columns) == 0{
				rSub.Columns = tb.GetColumnStr()
			}
		}

		if len(temp) != 0{
			for key,value := range temp{
				value.DbName = db.Name
				value.Rule = r
				value.SubRule = key
				job := Combination(value,param)
				if job == nil{
					continue
				}
				describe.Jobs = append(describe.Jobs, job)
			}
		}

	}


	return describe,nil
}

func IncrementDescribeFromSchema(
	sourceScheme *ConnectScheme,
	targetScheme *ConnectScheme,
	schema *Schema,
) (*Describe,error) {

	describe := &Describe{
		MySQL: sourceScheme.ToGoMysql(),
		Jobs:  make([]*Job, 0),
	}

	ruleC,err := NewRuleConfigs()


	if err != nil{
		return nil,err
	}

	if ruleC == nil{
		fmt.Println("解析rule规则出错")
		return nil,nil
	}


	for _,db := range schema.GetDBList(){
		r := ruleC.GetRule(db.Name)
		if r == nil{
			continue
		}

		tbs := db.GetTableList()

		temp := make(map[*RuleSub]*DataXContext,0)


		for _,tb := range tbs{

			rSub :=r.GetRuleSub(tb.Name)

			if rSub == nil{
				continue
			}

			sr,ok := temp[rSub]
			if !ok {
				sr = NewDataxContext()
				temp[rSub] = sr
			}
			sr.SourceTable = append(sr.SourceTable,tb.Name)
			// 重命名表
			sr.Sql = strings.Replace(tb.Sql,tb.Name,rSub.TargetTB,1)
			if rSub.Columns == nil || len(rSub.Columns) == 0{
				rSub.Columns = tb.GetColumnStr()
			}
		}

		if len(temp) != 0{
			for key,value := range temp{
				value.DbName = db.Name
				value.Rule = r
				value.SubRule = key
				job := Combination(value,nil)
				if job == nil{
					continue
				}
				describe.Jobs = append(describe.Jobs, job)
			}
		}

	}


	return describe,nil
}

type errMoreJobs struct {
	DB    string
	Table string
	Jobs  []*Job
}

type errNoCond struct {
	DB    string
	Table string
	Job   *Job
}

type errNotMatch struct {
	DB    string
	Table string
}

type errNotMatchJob struct {
	Job *Job
}

type DescribeErrors struct {
	Error        error
	MoreSync     []*errMoreJobs
	MoreUpdate   []*errMoreJobs
	NoCond       []*errNoCond
	NotMatch     []*errNotMatch
	NotMatchJobs []*errNotMatchJob
}

func (describe *Describe) ErrorsBySchema(schema *Schema) (errs *DescribeErrors) {

	type Match struct {
		Sync   []*Job
		Update []*Job
		Cond   bool
		DB     string
		Table  string
	}

	errs = &DescribeErrors{
		MoreSync:     make([]*errMoreJobs, 0),
		MoreUpdate:   make([]*errMoreJobs, 0),
		NoCond:       make([]*errNoCond, 0),
		NotMatch:     make([]*errNotMatch, 0),
		NotMatchJobs: make([]*errNotMatchJob, 0),
	}

	tables := make([]*Match, 0)

	schema.EachTable(func(db *DB, table *Table) error {

		tables = append(tables, &Match{
			DB:     db.Name,
			Table:  table.Name,
			Sync:   make([]*Job, 0),
			Update: make([]*Job, 0),
		})
		return nil
	})

	jobs, err := Complie(describe.Jobs)

	if err != nil {
		errs.Error = err
		return
	}

	for _, job := range jobs {

		find := false

		for _, table := range tables {

			if job.Sync.Matched(table.DB, table.Table) {
				table.Sync = append(table.Sync, job.Job)
				find = true
			}

			if job.Update.Matched(table.DB, table.Table) {
				table.Update = append(table.Update, job.Job)
				find = true
			}

		}

		if !find {
			errs.NotMatchJobs = append(errs.NotMatchJobs, &errNotMatchJob{
				Job: job.Job})
		}
	}

	for _, table := range tables {

		if len(table.Sync) > 1 {
			errs.MoreSync = append(errs.MoreSync, &errMoreJobs{
				DB:    table.DB,
				Table: table.Table,
				Jobs:  table.Sync,
			})

			continue
		}

		if len(table.Update) > 1 {
			errs.MoreSync = append(errs.MoreUpdate, &errMoreJobs{
				DB:    table.DB,
				Table: table.Table,
				Jobs:  table.Update,
			})

			continue
		}

		if len(table.Sync) == 0 && len(table.Update) == 0 {
			errs.NotMatch = append(errs.NotMatch, &errNotMatch{
				DB:    table.DB,
				Table: table.Table,
			})

			continue
		}

		for _, job := range table.Update {

			if !job.Update.IsHaveCond() {
				errs.NoCond = append(errs.NoCond, &errNoCond{
					Job:   job,
					DB:    table.DB,
					Table: table.Table,
				})
			}
		}
	}

	return
}

func (describe *Describe) Save(fileName string) error {

	bytes, err := json.Marshal(describe)

	if err != nil {
		return errors.ToFormatError(
			err,
			"marshal describe to json fail.",
		)
	}

	err = ioutil.WriteFile(fileName, bytes, 0755)

	if err != nil {
		return errors.ToFormatError(
			err,
			"write describe to file %s fail.", fileName,
		)
	}

	return nil
}
