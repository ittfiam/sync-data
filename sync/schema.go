package sync

import (
	"database/sql"

	"strings"

	"sync-mysql/errors"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Schema struct {
	Info      *SchemaInfo
	DataBases map[string]*DB
}

func (schema *Schema) EachTable(cb func(db *DB, table *Table) error) error {

	for _, db := range schema.GetDBList() {

		for _, table := range db.GetTableList() {

			err := cb(db, table)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (schema *Schema) GetDBList() []*DB {

	dbs := make([]*DB, len(schema.DataBases))

	for _, database := range schema.DataBases {
		dbs[database.Index] = database
	}

	return dbs

}

func (db *DB) GetTableList() []*Table {

	tables := make([]*Table, len(db.Tables))

	for _, table := range db.Tables {
		tables[table.Index] = table
	}

	return tables
}

type SchemaInfo struct {
	Skips  map[string]struct{}
	Prefix string
}

func (info *SchemaInfo) AddSkips(db ...string) {

	for _, name := range db {
		info.Skips[name] = struct{}{}
	}
}

func NewSchemaInfo() *SchemaInfo {

	info := &SchemaInfo{
		Skips: make(map[string]struct{}, 0),
	}

	return info

}

func NewSchemaFromMysql(mysql string, info *SchemaInfo) (schema *Schema, err error) {

	conn, err := sql.Open("mysql", mysql)

	defer conn.Close()

	if err != nil {
		err = errors.ToFormatError(
			err,
			"connect to mysql:(%s) fail.",
			mysql)

		return
	}

	schema, err = NewSchema(conn, info)

	return
}


func NewSchema(conn *sql.DB, info *SchemaInfo) (schema *Schema, err error) {


	if info == nil {
		info = NewSchemaInfo()
	}

	schema = &Schema{
		DataBases: make(map[string]*DB, 0),
		Info:      info,
	}

	rows, err := conn.Query(
		"select " +
			"TABLE_SCHEMA," +
			"TABLE_NAME," +
			"COLUMN_NAME," +
			"DATA_TYPE," +
			"COLUMN_KEY," +
			"COLUMN_COMMENT " +
			"from information_schema.COLUMNS")

	if err != nil {
		err = errors.ToFormatError(
			err,
			"select columns from information_schmea fail.",
		)

		return
	}

	var dbName string
	var tableName string
	var dbIndex int

	for rows.Next() {

		column := new(Column)

		err = rows.Scan(
			&dbName,
			&tableName,
			&column.Name,
			&column.Type,
			&column.Key,
			&column.Comment,
		)

		dbName = strings.TrimSpace(dbName)

		if err != nil {
			err = errors.ToFormatError(
				err,
				"scan row from information_schema fail.",
			)
			return
		}

		if _, ok := info.Skips[dbName]; ok ||
			dbName == "information_schema" ||
			dbName == "performance_schema" ||
			dbName == "mysql" {
			continue
		}

		if len(info.Prefix) > 0 && !strings.HasPrefix(dbName, info.Prefix) {
			continue
		}

		db, ok := schema.DataBases[dbName]

		if !ok {
			db = &DB{
				Name:   dbName,
				Tables: make(map[string]*Table, 0),
				Length: 0,
				Index:  dbIndex,
			}

			schema.DataBases[dbName] = db
			dbIndex++
		}

		table, ok := db.Tables[tableName]

		if !ok {
			table = NewTable(tableName, db.Length, db)
			db.Tables[tableName] = table
			db.Length++
		}

		table.append(column)
	}

	schema.EachTable(func(db *DB, table *Table) error {

		createSql := fmt.Sprintf("show create table %s.%s",db.Name,table.Name)
		rows := conn.QueryRow(createSql)

		var tableName string
		var tableSql string

		rows.Scan(&tableName,&tableSql)

		table.Sql = tableSql

		for index, column := range table.Columns {
			if column.Key == "PRI" {
				table.Keys = append(table.Keys, index)
			}
		}

		return nil
	})

	return
}
