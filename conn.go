// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package godbc

import (
	"context"
	"database/sql/driver"
	"strings"
	"unsafe"

	"github.com/zerobit-tech/godbc/api"
)

type Conn struct {
	h  api.SQLHDBC
	tx *Tx
}

func (d *Driver) Open(dsn string) (driver.Conn, error) {
	var out api.SQLHANDLE
	ret := api.SQLAllocHandle(api.SQL_HANDLE_DBC, api.SQLHANDLE(d.h), &out)
	if IsError(ret) {
		return nil, NewError("SQLAllocHandle", d.h)
	}
	h := api.SQLHDBC(out)
	drv.Stats.updateHandleCount(api.SQL_HANDLE_DBC, 1)

	b := api.StringToUTF16(dsn)
	ret = api.SQLDriverConnect(h, 0,
		(*api.SQLWCHAR)(unsafe.Pointer(&b[0])), api.SQLSMALLINT(len(b)),
		nil, 0, nil, api.SQL_DRIVER_NOPROMPT)
	if IsError(ret) {
		defer releaseHandle(h)
		return nil, NewError("SQLDriverConnect", h)
	}
	return &Conn{h: h}, nil
}

func (c *Conn) Close() error {
	ret := api.SQLDisconnect(c.h)
	if IsError(ret) {
		return NewError("SQLDisconnect", c.h)
	}
	h := c.h
	c.h = api.SQLHDBC(api.SQL_NULL_HDBC)
	return releaseHandle(h)
}

func (c *Conn) Ping(ctx context.Context) error {

	sqlToPing, ok := ctx.Value(SQL_TO_PING).(string)
	if !ok || strings.TrimSpace(sqlToPing) == "" {

		return nil

	}

	//fmt.Println("Pingging..:", sqlToPing)
	//args := make([]driver.Value, 0)

	var out api.SQLHANDLE
	//var os *ODBCStmt
	ret := api.SQLAllocHandle(api.SQL_HANDLE_STMT, api.SQLHANDLE(c.h), &out)
	if IsError(ret) {
		return NewError("SQLAllocHandle", c.h)
	}
	h := api.SQLHSTMT(out)
	//drv.Stats.updateHandleCount(api.SQL_HANDLE_STMT, 1)
	b := api.StringToUTF16(sqlToPing)
	ret = api.SQLExecDirect(h,
		(*api.SQLWCHAR)(unsafe.Pointer(&b[0])), api.SQL_NTS)
	if IsError(ret) {
		defer releaseHandle(h)
		return NewError("SQLExecDirectW", h)
	}
	// _, err := ExtractParameters(h)
	// if err != nil {
	// 	defer releaseHandle(h)
	// 	return err
	// }

	return nil

}

// Query method executes the statement with out prepare if no args provided, and a driver.ErrSkip otherwise (handled by sql.go to execute usual preparedStmt)
func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	if len(args) > 0 {
		// Not implemented for queries with parameters
		return nil, driver.ErrSkip
	}
	return c.QueryContext(context.Background(), query, make([]driver.NamedValue, 0))
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	dargs := make([]driver.Value, len(args))
	for n, param := range args {
		dargs[n] = param.Value
	}

	var out api.SQLHANDLE
	var os *ODBCStmt
	ret := api.SQLAllocHandle(api.SQL_HANDLE_STMT, api.SQLHANDLE(c.h), &out)
	if IsError(ret) {
		return nil, NewError("SQLAllocHandle", c.h)
	}
	h := api.SQLHSTMT(out)
	drv.Stats.updateHandleCount(api.SQL_HANDLE_STMT, 1)
	b := api.StringToUTF16(query)
	ret = api.SQLExecDirect(h,
		(*api.SQLWCHAR)(unsafe.Pointer(&b[0])), api.SQL_NTS)
	if IsError(ret) {
		defer releaseHandle(h)
		return nil, NewError("SQLExecDirectW", h)
	}
	ps, err := ExtractParameters(h)
	if err != nil {
		defer releaseHandle(h)
		return nil, err
	}
	os = &ODBCStmt{
		h:          h,
		Parameters: ps,
		usedByRows: true}
	err = os.BindColumns(ctx)
	if err != nil {
		return nil, err
	}
	return &Rows{os: os}, nil
}
