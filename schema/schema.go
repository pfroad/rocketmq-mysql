package schema

import (
	"github.com/golang/glog"
)

const (
	QUERY_SCHEMA = "select schema_name from schemata" +
		" where schema_name not in ('sys', 'performance_schema', 'mysql', 'information_schema')"
	QUERY_TABLE = "select table_name,column_name,data_type,column_type,character_set_name " +
		"from columns " +
		"where table_schema = ?"
)

type Schema struct {
	Name     string
	TableMap map[string]*Table
}

func (self *Schema) AddTable(table *Table) {
	tableName := table.Name
	if _, ok := self.TableMap[tableName]; !ok {
		self.TableMap[tableName] = table
	}
}

func (self *Schema) GetTable(tableName string) *Table {
	return self.TableMap[tableName]
}

var (
	SchemaMap map[string]*Schema
)

func Init() error {
	var schemaNames []string
	rows, err := GormDB.Raw(QUERY_SCHEMA).Rows()
	defer rows.Close()
	if err != nil {
		glog.Errorf("Query schema meta err: %v\n", err)
		return err
	}

	for rows.Next() {
		var schemaName string
		rows.Scan(&schemaName)
		schemaNames = append(schemaNames, schemaName)
	}

	SchemaMap = make(map[string]*Schema, len(schemaNames))
	for _, schemaName := range schemaNames {
		SchemaMap[schemaName] = loadSchema(schemaName)
	}

	return nil
}

type DBColumns struct {
	TableName   string
	ColumnName  string
	DataType    string
	ColumnType  string
	CharSetName string
}

func loadSchema(schemaName string) *Schema {
	var columns []DBColumns
	GormDB.Raw(QUERY_TABLE, schemaName).Scan(&columns)

	schema := &Schema{Name: schemaName, TableMap: make(map[string]*Table)}
	for _, column := range columns {
		table := schema.GetTable(column.TableName)
		if table == nil {
			table = &Table{SchemaName: schemaName, Name: column.TableName}
			schema.AddTable(table)
		}

		table.AddCol(column.ColumnName)
	}

	return schema
}
