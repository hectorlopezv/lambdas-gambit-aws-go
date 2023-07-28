package tools

import (
	"fmt"
	"strings"
	"time"
	"strconv"
)

func FechaMysql() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
func EscapeString(t string)string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}
func BuildQuery(s string, fieldName string, typeField string, valueN int, valueF float64, ValueS string)string{
	if typeField =="S"&& len(ValueS)== 0 ||
		typeField =="N"&& valueN== 0 ||
		typeField =="F"&& valueF== 0 {
		return s
	}
	if !strings.HasSuffix(s, "SET "){
		s += ", "
	}

	switch typeField {
	case "S":
		s += fieldName + " = '" + EscapeString(ValueS) + "'"
	case "N":
		s += fieldName + " = " + strconv.Itoa(valueN)
	case "F":
		s += fieldName + " = " + strconv.FormatFloat(valueF, 'e', -1, 64)
	}
	return s

}