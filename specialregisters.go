package godbc

import (
	"strings"
	"time"
)

func Dummy(x []byte) []byte {
	return nil
}
func CURRENT_DATE(x []byte) []byte {
	return []byte(time.Now().Local().Format("2006-01-02"))
}

func CURRENT_TIME(x []byte) []byte {
	return []byte(time.Now().Local().Format("2006-01-02"))
}

var DB2SpecialResigers map[string]func([]byte) []byte = map[string]func([]byte) []byte{
	"CURRENT CLIENT_ACCTNG":            Dummy,
	"CLIENT ACCTNG":                    Dummy,
	"CURRENT CLIENT_APPLNAME":          Dummy,
	"CLIENT APPLNAME":                  Dummy,
	"CURRENT CLIENT_PROGRAMID":         Dummy,
	"CLIENT PROGRAMID":                 Dummy,
	"CURRENT CLIENT_USERID":            Dummy,
	"CLIENT USERID":                    Dummy,
	"CURRENT CLIENT_WRKSTNNAME":        Dummy,
	"CLIENT WRKSTNNAME":                Dummy,
	"CURRENT DATE":                     Dummy,
	"CURRENT_DATE":                     Dummy,
	"CURRENT DEBUG MODE":               Dummy,
	"CURRENT DECFLOAT ROUNDING MODE":   Dummy,
	"CURRENT DEGREE":                   Dummy,
	"CURRENT IMPLICIT XMLPARSE OPTION": Dummy,
	"CURRENT PATH":                     Dummy,
	"CURRENT_PATH":                     Dummy,
	"CURRENT FUNCTION PATH":            Dummy,
	"CURRENT SCHEMA":                   Dummy,
	"CURRENT SERVER":                   Dummy,
	"CURRENT_SERVER":                   Dummy,
	"CURRENT TEMPORAL SYSTEM_TIME":     Dummy,
	"CURRENT TIME":                     Dummy,
	"CURRENT_TIME":                     Dummy,
	"CURRENT TIMESTAMP":                Dummy,
	"CURRENT_TIMESTAMP":                Dummy,
	"CURRENT TIMEZONE":                 Dummy,
	"CURRENT_TIMEZONE":                 Dummy,
	"CURRENT USER":                     Dummy,
	"CURRENT_USER":                     Dummy,
	"SESSION_USER":                     Dummy,
	"USER":                             Dummy,
	"SYSTEM_USER":                      Dummy,
}

func IsSepecialRegister(name string) bool {
	_, found := DB2SpecialResigers[strings.ToUpper(strings.TrimSpace(name))]
	return found
}

func GetSepecialValue(name string, param []byte) []byte {
	funcToCall, found := DB2SpecialResigers[strings.ToUpper(strings.TrimSpace(name))]
	if found {
		return funcToCall(param)
	}
	return nil
}
