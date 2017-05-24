package sync

import (
	"fmt"
	"sync-mysql/plugin/mysqlplugin"
	"sync-mysql/plugin/hdfsplugin"
	"errors"
)


func getColumns(dataxParam *DataXContext,param *CommandParam) ([]*hdfsplugin.Column,error){
	c := GetTransitionConfig()
	tc := c.GetTConfigItem(param.GetTransitionMode())
	if tc == nil{
		return nil,errors.New("未能获取列转换配置")
	}
	result := make([]*hdfsplugin.Column,0)
	for _,col := range dataxParam.Table.Columns{
		ts := tc.GetValue(col.Type)
		if ts == ""{
			return nil,errors.New(fmt.Sprintf("%s未能找到对应的live转换类型",col.Type))
		}
		item := &hdfsplugin.Column{
			Name:col.Name,
			Types:ts,
		}
		result = append(result,item)
	}

	return result,nil

}

/**
全量组合
 */
func HdfsCombinationInit(dataxParam *DataXContext,param *CommandParam) *Job{

	sourceScheme,err := param.GetSourceSchema()
	if err != nil{
		fmt.Println(err)
		return nil
	}


	reader := mysqlplugin.NewReader()
	writer := hdfsplugin.NewWriter()

	cr := mysqlplugin.NewConnectionReader()

	cr.JdbcUrl = append(cr.JdbcUrl, sourceScheme.ToDataXMysql(dataxParam.DbName))
	cr.Table = dataxParam.SourceTable
	reader.Parameter.Connection = append(reader.Parameter.Connection,cr)
	reader.Parameter.Username = sourceScheme.Username
	reader.Parameter.Password = sourceScheme.Password
	reader.Parameter.Column = dataxParam.SubRule.Columns

	cs,err := getColumns(dataxParam,param)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	writer.Parameter.DefaultFS = param.Target
	writer.Parameter.Path = fmt.Sprintf(param.Path,dataxParam.Rule.TargetDB +".db",dataxParam.SubRule.TargetTB)
	writer.Parameter.Column = cs


	work := NewWorker(reader,writer)

	job := new(Job)
	job.Name = fmt.Sprintf("%s.%s", dataxParam.DbName, dataxParam.SubRule.TargetTB)
	job.WRName = param.GetTransitionMode()
	job.Enable = true
	job.DB = dataxParam.DbName
	job.Collection = dataxParam.SubRule.TargetTB
	job.Work = work
	job.Sql = make([]string,0)
	job.Sql = append(job.Sql,writer.MakeCreateSql(dataxParam.Table.Name))
	return job
}

/**
增量组合
 */
func CombinationIncrement(dataxParam *DataXContext,param *CommandParam) *Job{


	sourceScheme,err := param.GetSourceSchema()
	if err != nil{
		fmt.Println(err)
		return nil
	}

	targetScheme,err := param.GetSourceSchema()
	if err != nil{
		fmt.Println(err)
		return nil
	}

	reader := mysqlplugin.NewReader()
	writer := mysqlplugin.NewWriter()

	cr := mysqlplugin.NewConnectionReader()
	cw := mysqlplugin.NewConnectionWriter()

	cr.JdbcUrl = append(cr.JdbcUrl, sourceScheme.ToDataXMysql(dataxParam.DbName))
	// 根据规则获取需要更新的表
	ts := dataxParam.SubRule.GetUpdateTable(dataxParam.SourceTable)
	if ts == nil || len(ts) <= 0{
		return nil
	}
	cr.Table = ts
	reader.Parameter.Connection = append(reader.Parameter.Connection,cr)
	reader.Parameter.Username = sourceScheme.Username
	reader.Parameter.Password = sourceScheme.Password
	reader.Parameter.Column = dataxParam.SubRule.Columns

	reader.Parameter.Where = dataxParam.SubRule.GetUpdateColumn() + " > '$last_update_date $last_update_time'"

	cw.JdbcUrl = targetScheme.ToDataXMysql("")
	cw.Table = append(cw.Table, dataxParam.SubRule.TargetTB)
	writer.Parameter.Connection = append(writer.Parameter.Connection, cw)
	writer.Parameter.Username = targetScheme.Username
	writer.Parameter.Password = targetScheme.Password
	writer.Parameter.Column = dataxParam.SubRule.Columns

	writer.Parameter.WriteMode = "replace"


	work := NewWorker(reader,writer)

	job := new(Job)
	job.Name = fmt.Sprintf("%s.%s", dataxParam.DbName, dataxParam.SubRule.TargetTB)
	job.Enable = true
	job.DB = dataxParam.DbName
	job.Collection = dataxParam.SubRule.TargetTB
	job.Work = work
	return job
}