package Utilities

import (
	"database/sql"
	"strconv"
	"time"
)

func GetNullTime() time.Time {
	return time.Time{}
}

func GenNullInt(i int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(i), Valid: i > 0}
}

func GenNullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: str != ""}
}

func TimeToSqlString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}

func TimeToJsonString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00")
}

func JsonStringToTime(str string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z07:00", str)
	return t
}

func StringToTime(str string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05.000", str)
	return t
}

func StringToInt(str string) int {
	int, _ := strconv.Atoi(str)
	return int
}

func StringToFloat64(str string) float64 {
	float, _ := strconv.ParseFloat(str, 64)
	return float
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func StringToBool(str string) bool {
	if str == "true" {
		return true
	}
	return false
}

func StringsToInts(strings []string) []int {
	ints := []int{}
	for i := 0; i < len(strings); i++ {
		integer := StringToInt(strings[i])
		ints = append(ints, integer)
	}
	return ints
}

func IntsToStrings(ints []int) []string {
	strings := []string{}
	for i := 0; i < len(ints); i++ {
		str := IntToString(ints[i])
		strings = append(strings, str)
	}
	return strings
}

func BoolsToStrings(bools []bool) []string {
	strings := []string{}
	for i := 0; i < len(bools); i++ {
		str := BoolToString(bools[i])
		strings = append(strings, str)
	}
	return strings
}

func StringsToJsonObjectArray(strings []string) string {
	str := "["
	for i, s := range strings {
		str += s
		if i < len(strings)-1 {
			str += ","
		}
	}
	str += "]"
	return str
}

func StringsToJsonStringArray(strings []string) string {
	str := "["
	for i, s := range strings {
		str += "\"" + s + "\""
		if i < len(strings)-1 {
			str += ","
		}
	}
	str += "]"
	return str
}

func IntToFloat64(i int) float64 {
	return float64(i)
}

func GenNaturalIntSlice(length int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = i + 1
	}
	return numbers
}

func GenConstantIntSlice(length int, constant int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = constant
	}
	return numbers
}

func GenConstantStringSlice(length int, constant string) []string {
	strings := make([]string, length)
	for i := 0; i < length; i++ {
		strings[i] = constant
	}
	return strings
}

func GenConstantBoolSlice(length int, constant bool) []bool {
	bools := make([]bool, length)
	for i := 0; i < length; i++ {
		bools[i] = constant
	}
	return bools
}
