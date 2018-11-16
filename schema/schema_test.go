package schema

import (
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestLoadSchema(t *testing.T) {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return defaultTableName[:len(defaultTableName)-1]
	// }

	db := "apuser:airparking@tcp(10.35.22.61:3306)/information_schema?parseTime=true"
	var err error
	GormDB, err = gorm.Open("mysql", db)
	if err != nil {
		t.Errorf("Failed to connect DB %s, %v", db, err)
	}

	var schemaNames []string
	// GormDB.
	// 	Table("schemata").
	// 	Select("schema_name").Where("schema_name not in ('sys', 'performance_schema', 'mysql', 'information_schema')").
	// 	Scan(&schemaNames)
	rows, err := GormDB.Raw(QUERY_SCHEMA).Rows()
	defer rows.Close()
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var schemaName string
		rows.Scan(&schemaName)
		schemaNames = append(schemaNames, schemaName)
	}
	schemaMap := make(map[string]*Schema, len(schemaNames))
	for _, schemaName := range schemaNames {
		schemaMap[schemaName] = loadSchema(schemaName)
	}

	for k, v := range schemaMap {
		if k == "airparking" {
			table := v.GetTable("ap_user")
			columns := table.ColNames

			for i, column := range columns {
				log.Printf("%d-%s", i, column)
			}
		}
	}
}
