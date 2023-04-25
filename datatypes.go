package godbc

var SPParamDataTypes map[string]string = map[string]string{

	"SMALLINT":         "NUMERIC",
	"BIGINT":           "NUMERIC",
	"ROWID":            "NUMERIC",
	"DECIMAL":          "NUMERIC",
	"NUMERIC":          "NUMERIC",
	"INTEGER":          "NUMERIC",
	"DOUBLE PRECISION": "NUMERIC",
	"REAL":             "NUMERIC",

	"CHARACTER":                          "STRING",
	"BINARY VARYING":                     "STRING",
	"XML":                                "STRING",
	"BINARY LARGE OBJECT":                "STRING",
	"GRAPHIC VARYING":                    "STRING",
	"CHARACTER VARYING":                  "STRING",
	"BINARY":                             "STRING",
	"CHARACTER LARGE OBJECT":             "STRING",
	"GRAPHIC":                            "STRING",
	"DOUBLE-BYTE CHARACTER LARGE OBJECT": "STRING",

	"TIME":      "DATE",
	"DATE":      "DATE",
	"TIMESTAMP": "DATE",

	//"DISTINCT"
	//"ARRAY"

}

var SPParamStringTypes map[string]string = map[string]string{

	"CHARACTER":                          "STRING",
	"BINARY VARYING":                     "STRING",
	"XML":                                "STRING",
	"BINARY LARGE OBJECT":                "STRING",
	"GRAPHIC VARYING":                    "STRING",
	"CHARACTER VARYING":                  "STRING",
	"BINARY":                             "STRING",
	"CHARACTER LARGE OBJECT":             "STRING",
	"GRAPHIC":                            "STRING",
	"DOUBLE-BYTE CHARACTER LARGE OBJECT": "STRING",
}

var SPParamNumericTypes map[string]string = map[string]string{

	"SMALLINT":         "NUMERIC",
	"BIGINT":           "NUMERIC",
	"ROWID":            "NUMERIC",
	"DECIMAL":          "NUMERIC",
	"NUMERIC":          "NUMERIC",
	"INTEGER":          "NUMERIC",
	"DOUBLE PRECISION": "NUMERIC",
	"REAL":             "NUMERIC",
}

// without decimal point
var SPParamIntegerTypes map[string]string = map[string]string{

	"SMALLINT": "NUMERIC",
	"BIGINT":   "NUMERIC",
	"INTEGER":  "NUMERIC",
}

var SPParamDateTypes map[string]string = map[string]string{

	"TIME":      "DATE",
	"DATE":      "DATE",
	"TIMESTAMP": "DATE",

	//"DISTINCT"
	//"ARRAY"

}
