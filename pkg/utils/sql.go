package utils

import "database/sql"

func Str2NullStr(str string) (nullStr sql.NullString) {
	if str == "" {
		nullStr.String = ""
		nullStr.Valid = false
	} else {
		nullStr.String = str
		nullStr.Valid = true
	}
	return
}

func NullStr2Str(str sql.NullString) (nullStr string) {
	if str.Valid {
		nullStr = str.String
	} else {
		nullStr = ""
	}
	return
}
