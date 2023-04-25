package godbc

type ROWID int64
type ContextKey string

const LOAD_SP_RESULT_SETS ContextKey = "LOAD_SP_RESULT_SETS"
const DUMMY_SP_CALL ContextKey = "DUMMY_SP_CALL"
const ESCAPE_QUOTE ContextKey = "ESCAPE_QUOTE"
