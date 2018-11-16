package schema

import (
	"encoding/base64"
	"regexp"
	"strings"
)

type CloumnParser interface {
	GetValue(value interface{}) interface{}
}

func extractEnumValues(colType string) []string {
	var enumValues []string
	reg := regexp.MustCompile("(enum|set)\\((.*)\\)")
	if reg.MatchString(colType) {
		enumValues = strings.Split(strings.Replace(reg.ReplaceAllString(colType, "$2"), "'", "", -1), ",")
		// enumValues = strings.Split(reg.ReplaceAllString(colType, "$2"), ",")
	}

	return enumValues
}

func GetColumnParser(dataType, colType, charset string) CloumnParser {
	switch dataType {
	case "tinyint", "smallint", "mediumint", "int":
		// return;
	case "bigint":
	case "tinytext":
	case "text":
	default:
		return &DefaultColumnParser{}
	}

	return nil
}

type DefaultColumnParser struct {
}

func (self *DefaultColumnParser) GetValue(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case []byte:
		return base64.StdEncoding.EncodeToString(value.([]byte))
	}

	return value
}
