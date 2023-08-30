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

	"CURRENT ACCELERATOR":                    Dummy,
	"CURRENT APPLICATION COMPATIBILITY":      Dummy,
	"CURRENT APPLICATION ENCODING SCHEME":    Dummy,
	"CURRENT CLIENT_CORR_TOKEN":              Dummy,
	"CURRENT EXPLAIN MODE":                   Dummy,
	"CURRENT GET_ACCEL_ARCHIVE":              Dummy,
	"CURRENT_LC_CTYPE":                       Dummy,
	"CURRENT LOCALE LC_CTYPE":                Dummy,
	"CURRENT MAINTAINED":                     Dummy,
	"CURRENT MEMBER":                         Dummy,
	"CURRENT OPTIMIZATION HINT":              Dummy,
	"CURRENT PACKAGE PATH":                   Dummy,
	"CURRENT PACKAGESET":                     Dummy,
	"CURRENT PRECISION":                      Dummy,
	"CURRENT QUERY ACCELERATION":             Dummy,
	"CURRENT QUERY ACCELERATION WAITFORDATA": Dummy,
	"CURRENT REFRESH AGE":                    Dummy,
	"CURRENT ROUTINE VERSION":                Dummy,
	"CURRENT RULES":                          Dummy,
	"CURRENT_SCHEMA":                         Dummy,
	"CURRENT SQLID":                          Dummy,
	"CURRENT TEMPORAL BUSINESS_TIME":         Dummy,
	"CURRENT TIME ZONE":                      Dummy,
	"ENCRYPTION PASSWORD":                    Dummy,
	"SESSION TIME ZONE":                      Dummy,
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
