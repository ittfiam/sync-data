package sync

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"encoding/json"

	"bytes"

	"encoding/binary"

	"crypto/sha256"

	"reflect"

	"sync-mysql/errors"
)

type Table struct {
	Name    string
	DB      *DB
	Index   int
	Columns []*Column
	Keys    []int
	buffer  *bytes.Buffer
	Sql string
}

func (table *Table) GetSQL() string {

	columns := make([]string, 0)

	for _, column := range table.Columns {
		columns = append(columns, column.Name)
	}

	c := strings.Join(columns, ",")

	return fmt.Sprintf("select %s from %s.%s", c, table.DB.Name, table.Name)
}

func (table *Table) Scan(index int, rows *sql.Rows) (doc map[string]interface{}, key string, err error) {

	raw := table.newRecord()
	record := make([]interface{}, len(table.Columns))

	err = rows.Scan(raw...)

	if err != nil {
		err = errors.ToFormatError(
			err,
			"scan rows fail.",
		)

		return
	}

	for index, value := range raw {

		if value == nil {
			record[index] = nil
		} else {
			record[index] = reflect.ValueOf(value).Elem().Interface()
		}

	}

	err = table.recordFormat(index, record)

	if err != nil {
		return
	}

	key, err = table.recordKey(record)

	if err != nil {
		return
	}

	doc = table.toDocument(record)

	return

}

func (table *Table) newRecord() (raw []interface{}) {

	size := len(table.Columns)
	raw = make([]interface{}, size)

	for index, column := range table.Columns {
		switch column.Type {
		case "smallint", "int", "bigint", "tinyint", "enum":
			raw[index] = new(int64)
		case "datetime", "timestamp", "date":
			raw[index] = new(time.Time)
		case "varchar", "text", "decimal", "char", "mediumtext":
			raw[index] = new(string)
		case "blob":
			raw[index] = new([]byte)
		case "float", "double":
			raw[index] = new(float64)

		default:
			panic(
				fmt.Sprintf("sql type<%s> not support", column.Type))
		}
	}

	return
}
func (table *Table) recordKey(record []interface{}) (key string, err error) {

	table.buffer.Reset()

	if len(table.Keys) == 0 {
		err = errors.NewError(
			"key not found in %s.%s", table.DB.Name, table.Name,
		)

		return
	}

	if len(table.Keys) == 1 {

		key = fmt.Sprintf("%v", record[table.Keys[0]])
		return
	}

	for _, index := range table.Keys {

		var value interface{} = nil

		raw := record[index]

		if raw == nil {
			raw = "nil"
		}

		switch node := raw.(type) {
		case string:
			value = []byte(node)
		case time.Time:
			value, err = node.MarshalBinary()

			if err != nil {
				value = node.Unix()
			}
		default:
			value = raw
		}

		err = binary.Write(table.buffer, binary.BigEndian, value)

		if err != nil {
			err = errors.ToFormatError(
				err,
				"calc keys for %s.%s fail.", table.DB.Name, table.Name,
			)
			return
		}
	}

	key = fmt.Sprintf("%x", sha256.Sum256(table.buffer.Bytes()))

	return
}
func (table *Table) toDocument(record []interface{}) (doc map[string]interface{}) {

	doc = make(map[string]interface{}, 0)

	for index, column := range table.Columns {
		doc[column.Name] = record[index]
	}

	return
}

func (table *Table) recordFormat(index int, record []interface{}) (err error) {

	for index, value := range record {

		switch node := value.(type) {
		case string:

			if len(node) == 0 {
				value = node
				break
			}

			var inner interface{}

			switch (node)[0] {
			case '{':
				inner = make(map[string]interface{}, 0)
			case '[':
				inner = make([]interface{}, 0)
			}

			if inner != nil {
				err := json.Unmarshal([]byte(node), &inner)

				if err != nil {
					fmt.Printf("WARNING:%s in %s.%s : %v\n", err.Error(), table.DB.Name, table.Name, index)
				} else {
					record[index] = inner
				}
			}
		}
	}

	return
}

func (table *Table) ColumnExists(name string) bool {

	for _, column := range table.Columns {

		if column.Name == name {
			return true
		}
	}

	return false
}

func (table *Table) GetColumnStr() []string{
	result := make([]string,0)
	for _, column := range table.Columns {
		result = append(result,column.Name)
	}
	return result
}

func NewTable(name string, index int, db *DB) *Table {
	table := &Table{
		Name:    name,
		Columns: make([]*Column, 0),
		Index:   index,
		DB:      db,
		Keys:    make([]int, 0),
		buffer:  bytes.NewBuffer(nil),
	}

	return table
}

func (table *Table) append(column ...*Column) {
	table.Columns = append(table.Columns, column...)
}
