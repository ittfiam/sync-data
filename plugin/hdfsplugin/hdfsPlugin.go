package hdfsplugin

import (
	"fmt"
	"strings"
)

type Column struct {

	Name string	`json:"name"`
	Types string	`json:"type"`
}

func (c *Column) ToHqlString() string{

	return fmt.Sprintf("%s %s",c.Name,c.Types)
}

type Parameter struct {

	DefaultFS string	`json:"defaultFS"`
	FileType string		`json:"fileType"`
	Path	string		`json:"path"`
	FileName string		`json:"fileName"`
	Column []*Column	`json:"column"`
	WriteMode string 	`json:"writeMode"`
	FieldDelimiter string	`json:"fieldDelimiter"`
}

type Writer struct {
	Name string	`json:"name"`
	Parameter *Parameter	`json:"parameter"`
}


func NewParameter() *Parameter{

	return &Parameter{

		DefaultFS:"",
		FileType:"text",
		Path:"",
		FileName:"",
		Column:make([]*Column,0),
		WriteMode:"append",
		FieldDelimiter:"\t",

	}
}

func NewWriter() *Writer{

	return &Writer{
		Name:"hdfswriter",
		Parameter:NewParameter(),
	}
}


func (w *Writer) MakeCreateSql(tableName string) string{

	s := make([]string,0)
	for _,c := range w.Parameter.Column{
		s = append(s,c.ToHqlString())
	}
	rows := strings.Join(s,",")

	return fmt.Sprintf(
		"create table %s(%s) row format delimited fields terminated by \"\t\" STORED AS TEXTFILE;",
		tableName,
		rows,
	)
}
