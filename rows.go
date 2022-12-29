// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package godbc

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/zerobit-tech/godbc/database/sql/driver"

	"github.com/zerobit-tech/godbc/api"
)

type Rows struct {
	os *ODBCStmt
}

func (r *Rows) Columns() []string {
	names := make([]string, len(r.os.Cols))
	for i := 0; i < len(names); i++ {
		names[i] = r.os.Cols[i].Name()
	}
	return names
}

func (r *Rows) ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) {
	//TODO(Akhil):This functions retuns the precision and scale of column.
	ok = false
	var namelen api.SQLSMALLINT
	namebuf := make([]byte, api.MAX_FIELD_SIZE)
	buf := api.SQLLEN(32)
	ret := api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_TYPE_NAME, api.SQLPOINTER(unsafe.Pointer(&namebuf[0])), (api.MAX_FIELD_SIZE), (*api.SQLSMALLINT)(&namelen), &buf)

	if IsError(ret) {
		fmt.Println(ret)
		return 0, 0, false
	}

	dbtype := string(namebuf[:namelen])
	buf2 := api.SQLLEN(32)
	ret = api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_PRECISION, api.SQLPOINTER(unsafe.Pointer(nil)), 0, (*api.SQLSMALLINT)(nil), &buf2)
	if IsError(ret) {
		fmt.Println(ret)
		return 0, 0, false
	}

	precision = int64(buf2)

	buf3 := api.SQLLEN(32)

	ret = api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_SCALE, api.SQLPOINTER(unsafe.Pointer(nil)), 0, (*api.SQLSMALLINT)(nil), &buf3)

	scale = int64(buf3)

	if IsError(ret) {
		fmt.Println(ret)
		return 0, 0, false
	}
	if dbtype == "DECIMAL" {
		ok = true
	} else if dbtype == "NUMERIC" {
		ok = true
	} else if dbtype == "TIMESTAMP" {
		ok = true
	}
	return precision, scale, ok
}

func (r *Rows) ColumnTypeLength(index int) (length int64, ok bool) {
	//ToDo(Akhil):This functions retuns the length of column.
	buf := api.SQLLEN(32)
	ret := api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_LENGTH, api.SQLPOINTER(unsafe.Pointer(nil)), 0, (*api.SQLSMALLINT)(nil), &buf)
	length = int64(buf)
	if IsError(ret) {
		fmt.Println(ret)
		return 0, false
	}
	return length, true
}

func (r *Rows) ColumnTypeScanType(index int) reflect.Type {
	//TODO(AKHIL):This function will return the scantype that can be used to scan
	//the data to the golang variable.
	a := r.os.Cols[index].TypeScan()
	return (a)
}

func (r *Rows) ColumnTypeNullable(index int) (nullable, ok bool) {
	//TODO(Akhil):This functions retuns whether the column is nullable or not
	var null int64
	buf := api.SQLLEN(32)
	ret := api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_NULLABLE, api.SQLPOINTER(unsafe.Pointer(nil)), 0, (*api.SQLSMALLINT)(nil), &buf)
	if IsError(ret) {
		fmt.Println(ret)
		return false, false
	}
	if null == api.SQL_NULLABLE {
		return true, true
	}
	return false, true
}

func (r *Rows) ColumnTypeDatabaseTypeName(index int) string {
	//TODO(AKHIL):This functions retuns the dbtype(VARCHAR,DECIMAL etc..) of column.
	//namebuf can be of uint8 or byte
	var namelen api.SQLSMALLINT
	namebuf := make([]byte, api.MAX_FIELD_SIZE)
	buf := api.SQLLEN(32)
	ret := api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_TYPE_NAME, api.SQLPOINTER(unsafe.Pointer(&namebuf[0])), (api.MAX_FIELD_SIZE), (*api.SQLSMALLINT)(&namelen), &buf)

	if IsError(ret) {
		fmt.Println(ret)
		return ""
	}
	dbtype := string(namebuf[:namelen])
	return dbtype
}

func (r *Rows) ColumnTypeLabel(index int) string {

	var namelen api.SQLSMALLINT
	namebuf := make([]byte, api.MAX_FIELD_SIZE)
	buf := api.SQLLEN(32)
	ret := api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_LABEL, api.SQLPOINTER(unsafe.Pointer(&namebuf[0])), (api.MAX_FIELD_SIZE), (*api.SQLSMALLINT)(&namelen), &buf)

	if IsError(ret) {
		fmt.Println(ret)
		return ""
	}
	dbtype := string(namebuf[:namelen])
	return dbtype
}

func (r *Rows) ColumnTypeTable(index int) string {

	var namelen api.SQLSMALLINT
	namebuf := make([]byte, api.MAX_FIELD_SIZE)
	buf := api.SQLLEN(32)
	ret := api.SQLColAttribute(r.os.h, api.SQLUSMALLINT(index+1), api.SQL_DESC_TABLE_NAME, api.SQLPOINTER(unsafe.Pointer(&namebuf[0])), (api.MAX_FIELD_SIZE), (*api.SQLSMALLINT)(&namelen), &buf)

	if IsError(ret) {
		fmt.Println(ret)
		return ""
	}
	dbtype := string(namebuf[:namelen])
	return dbtype
}

func (r *Rows) Next(dest []driver.Value) error {
	ret := api.SQLFetch(r.os.h)
	if ret == api.SQL_NO_DATA {
		return io.EOF
	}
	if IsError(ret) {
		return NewError("SQLFetch", r.os.h)
	}
	for i := range dest {
		v, err := r.os.Cols[i].Value(r.os.h, i)

		if err != nil {
			fmt.Println("dest[i] = v value err>>>", dest[i], err.Error())
			v = fmt.Sprintf("**ERROR** : %s", err.Error())
			//return err
		}

		dest[i] = v
	}
	return nil
}

func (r *Rows) JumpToRow(rowNumber int) error {
	row := api.SQLSETPOSIROW(rowNumber)
	ret := api.SQLSetPos(r.os.h, row, api.SQL_POSITION, api.SQL_LOCK_NO_CHANGE)
	if IsError(ret) {
		return NewError("SQLSetPos", r.os.h)
	}
	return nil
}

func (r *Rows) JumpToRow2(rowNumber int) error {
	rowsToMove := api.SQLLEN(rowNumber)
	ret := api.SQLFetchScroll(r.os.h, api.SQL_FETCH_ABSOLUTE, rowsToMove)
	if IsError(ret) {
		return NewError("SQLSetPos", r.os.h)
	}
	return nil
}

func (r *Rows) Close() error {
	return r.os.closeByRows()
}

func (r *Rows) HasNextResultSet() bool {
	return true
}

func (r *Rows) NextResultSet() error {
	ret := api.SQLMoreResults(r.os.h)
	if ret == api.SQL_NO_DATA {
		return io.EOF
	}
	if IsError(ret) {
		return NewError("SQLMoreResults", r.os.h)
	}

	err := r.os.BindColumns()
	if err != nil {
		return err
	}
	return nil
}
