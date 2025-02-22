package sync

import (
	"fmt"
	"sync-mysql/plugin/mysqlplugin"
)

/**
全量组合
 */

func MysqlMappingMode(dataxParam *DataXContext,param *CommandParam) *Job{

	if param.Mode == 0{
		return MysqlCombinationInit(dataxParam,param)
	} else if param.Mode == 1{
		return MysqlCombinationIncrement(dataxParam,param)
	}

	return nil
}

func MysqlCombinationInit(dataxParam *DataXContext,param *CommandParam) *Job{

	sourceScheme,err := param.GetSourceSchema()
	if err != nil{
		fmt.Println(err)
		return nil
	}

	targetScheme,err := param.GetTargetSchema()
	if err != nil{
		fmt.Println(err)
		return nil
	}

	reader := mysqlplugin.NewReader()
	writer := mysqlplugin.NewWriter()

	cr := mysqlplugin.NewConnectionReader()
	cw := mysqlplugin.NewConnectionWriter()

	cr.JdbcUrl = append(cr.JdbcUrl, sourceScheme.ToDataXMysql(dataxParam.DbName))
	cr.Table = dataxParam.SourceTable
	reader.Parameter.Connection = append(reader.Parameter.Connection,cr)
	reader.Parameter.Username = sourceScheme.Username
	reader.Parameter.Password = sourceScheme.Password
	reader.Parameter.Column = dataxParam.SubRule.Columns

	cw.JdbcUrl = targetScheme.ToDataXMysql("")
	cw.Table = append(cw.Table, dataxParam.SubRule.TargetTB)
	writer.Parameter.Connection = append(writer.Parameter.Connection, cw)
	writer.Parameter.Username = targetScheme.Username
	writer.Parameter.Password = targetScheme.Password
	writer.Parameter.Column = dataxParam.SubRule.Columns

	if !dataxParam.SubRule.NotNeedTruncate{
		dropSql := writer.MakeDeleteSql(dataxParam.SubRule.TargetTB)
		writer.Parameter.PreSql = append(writer.Parameter.PreSql,dropSql)
	}
	writer.Parameter.WriteMode = "insert"



	work := NewWorker(reader,writer)

	job := new(Job)
	job.Name = fmt.Sprintf("%s.%s", dataxParam.DbName, dataxParam.SubRule.TargetTB)
	job.WRName = param.GetTransitionMode()
	job.Enable = true
	job.DB = dataxParam.DbName
	job.Collection = dataxParam.SubRule.TargetTB
	job.Work = work
	job.Sql = make([]string,0)
	job.Sql = append(job.Sql,writer.MakeDropSql(dataxParam.SubRule.TargetTB))
	job.Sql = append(job.Sql,writer.MakeCreateSql(dataxParam.Sql))
	return job
}

/**
增量组合
 */
func MysqlCombinationIncrement(dataxParam *DataXContext,param *CommandParam) *Job{


	sourceScheme,err := param.GetSourceSchema()
	if err != nil{
		fmt.Println(err)
		return nil
	}

	targetScheme,err := param.GetTargetSchema()
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
	job.WRName = param.GetTransitionMode()
	job.Enable = true
	job.DB = dataxParam.DbName
	job.Collection = dataxParam.SubRule.TargetTB
	job.Work = work
	return job
}