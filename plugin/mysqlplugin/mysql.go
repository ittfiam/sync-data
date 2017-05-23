package mysqlplugin

import (
	"fmt"
	"strings"
)

type ConnectionReader struct {
	JdbcUrl []string `json:"jdbcUrl"`
	Table []string `json:"table"`
}

type ConnectionWriter struct {
	JdbcUrl string `json:"jdbcUrl"`
	Table []string `json:"table"`
}


func NewConnectionReader() *ConnectionReader {

	return &ConnectionReader{
		JdbcUrl:make([]string,0),
		Table:make([]string,0),
	}
}


func NewConnectionWriter() *ConnectionWriter {

	return &ConnectionWriter{
		JdbcUrl:"",
		Table:make([]string,0),
	}
}



type ReadParameter struct {
	Column []string `json:"column"`
	Connection []*ConnectionReader `json:"connection"`
	Username string `json:"username"`
	Password string `json:"password"`
	Where string `json:"where"`
}

func newReadParameter() *ReadParameter{

	return &ReadParameter{
		Column:make([]string,0),
		Connection:make([]*ConnectionReader,0),
		Username:"",
		Password:"",
		Where:"",
	}
}

type Reader struct {
	Name string `json:"name"`
	Parameter *ReadParameter `json:"parameter"`
}


func NewReader() *Reader{

	return &Reader{
		Name:"mysqlreader",
		Parameter:newReadParameter(),
	}
}



type WriteParameter struct {
	Column []string `json:"column"`
	Connection []*ConnectionWriter `json:"connection"`
	PreSql []string `json:"preSql"`
	Session []string `json:"session"`
	Username string `json:"username"`
	Password string `json:"password"`
	WriteMode string `json:"writeMode"`
}

func newWriteParameter() *WriteParameter{

	return &WriteParameter{
		Column:     make([]string,0),
		Connection: make([]*ConnectionWriter,0),
		PreSql:     make([]string,0),
		Session:    make([]string,0),
		Username:   "",
		Password:   "",
		WriteMode:  "",
	}

}

type Writer struct{
	Name string `json:"name"`
	Parameter *WriteParameter `json:"parameter"`
}

func NewWriter() *Writer{

	return &Writer{
		Name:      "mysqlwriter",
		Parameter: newWriteParameter(),
	}
}


func (writer *Writer) MakeDropSql(tableName string) string{

	return fmt.Sprintf("drop table if exists %s;",tableName)
}

func (writer *Writer) MakeDeleteSql(tableName string) string{

	return fmt.Sprintf(" TRUNCATE TABLE %s;",tableName)
}

func (writer *Writer) MakeCreateSql(createTableSql string) string{

	return strings.Replace(createTableSql,"NOT NULL","",-1)
}


