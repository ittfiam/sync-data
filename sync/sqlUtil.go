package sync

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync-mysql/errors"
)

func ExecSql(cn string,sqls... string) error{

	conn, err := sql.Open("mysql", cn)

	defer conn.Close()

	if err != nil {
		err = errors.ToFormatError(
			err,
			"connect to mysql:(%s) fail.",
			cn)

		return err
	}

	for _,s:=range sqls{

		if s == ""{
			continue
		}
		_,err := conn.Exec(s)

		if err != nil{
			return  errors.ToFormatError(
				err,
				"execute sql fail sql:(%s) .",
				s)
		}
	}

	return nil

}
