package schema

import (
	"testing"
)

func TestAxtractEnumValues(t *testing.T) {
	values := extractEnumValues("enum('F','M','C')")

	// test := []byte{1, 2}

	// typeName := typeof(test)
	// typeName := fmt.Sprint(reflect.TypeOf(test))
	// t.Log(typeName)
	table := &Table{SchemaName: "db1", Name: "sbtest1"}
	table.AddCol("a")
	for _, value := range values {
		if value != "F" && value != "M" && value != "C" {
			t.Fail()
		}
	}
}

func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case []byte:
		return "[]byte"
	default:
		_ = t
		return "unknown"
	}
}
