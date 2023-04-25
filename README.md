# godbc

Interface for GoLang to DB2 for z/OS, DB2 for LUW, DB2 for i.
I merge code from following two repos
	https://github.com/alexbrainman/odbc
	https://github.com/ibmdb/go_ibm_db



## Prerequisite

Golang should be installed(Golang version should be >=1.12.x and <= 1.19.X)

Git should be installed in your system.

For non-windows users, GCC and tar should be present in your system.

```
For Docker Linux Container(Ex: Amazon Linux2), use below commands:
yum install go git tar libpam
```
 

## How to run sample program

### example1.go:-

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/onlysumitg/GoQhttp"
)

func main() {
	con := "HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	db, err := sql.Open("godbc", con)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
```
To run the sample:- go run example1.go

For complete list of connection parameters please check [this.](https://www.ibm.com/support/knowledgecenter/SSEPGG_11.1.0/com.ibm.swg.im.dbclient.config.doc/doc/c0054698.html)

### example2.go:-

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/onlysumitg/GoQhttp"
)

func Create_Con(con string) *sql.DB {
	db, err := sql.Open("godbc", con)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

// Creating a table.

func create(db *sql.DB) error {
	_, err := db.Exec("DROP table SAMPLE")
	if err != nil {
		_, err := db.Exec("create table SAMPLE(ID varchar(20),NAME varchar(20),LOCATION varchar(20),POSITION varchar(20))")
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec("create table SAMPLE(ID varchar(20),NAME varchar(20),LOCATION varchar(20),POSITION varchar(20))")
		if err != nil {
			return err
		}
	}
	fmt.Println("TABLE CREATED")
	return nil
}

// Inserting row.

func insert(db *sql.DB) error {
	st, err := db.Prepare("Insert into SAMPLE(ID,NAME,LOCATION,POSITION) values('3242','mike','hyd','manager')")
	if err != nil {
		return err
	}
	st.Query()
	return nil
}

// This API selects the data from the table and prints it.

func display(db *sql.DB) error {
	st, err := db.Prepare("select * from SAMPLE")
	if err != nil {
		return err
	}
	err = execquery(st)
	if err != nil {
		return err
	}
	return nil
}

func execquery(st *sql.Stmt) error {
	rows, err := st.Query()
	if err != nil {
		return err
	}
	cols, _ := rows.Columns()
	fmt.Printf("%s    %s   %s    %s\n", cols[0], cols[1], cols[2], cols[3])
	fmt.Println("-------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var t, x, m, n string
		err = rows.Scan(&t, &x, &m, &n)
		if err != nil {
			return err
		}
		fmt.Printf("%v  %v   %v         %v\n", t, x, m, n)
	}
	return nil
}

func main() {
	con := "HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	type Db *sql.DB
	var re Db
	re = Create_Con(con)
	err := create(re)
	if err != nil {
		fmt.Println(err)
	}
	err = insert(re)
	if err != nil {
		fmt.Println(err)
	}
	err = display(re)
	if err != nil {
		fmt.Println(err)
	}
}
```
To run the sample:- go run example2.go

### example3.go:-(POOLING)

```go
package main

import (
	_ "database/sql"
	"fmt"
	a "github.com/onlysumitg/GoQhttp"
)

func main() {
	con := "HOSTNAME=host;PORT=number;DATABASE=name;UID=username;PWD=password"
	pool := a.Pconnect("PoolSize=100")

	// SetConnMaxLifetime will take the value in SECONDS
	db := pool.Open(con, "SetConnMaxLifetime=30")
	st, err := db.Prepare("Insert into SAMPLE values('hi','hi','hi','hi')")
	if err != nil {
		fmt.Println(err)
	}
	st.Query()

	// Here the time out is default.
	db1 := pool.Open(con)
	st1, err := db1.Prepare("Insert into SAMPLE values('hi1','hi1','hi1','hi1')")
	if err != nil {
		fmt.Println(err)
	}
	st1.Query()

	db1.Close()
	db.Close()
	pool.Release()
	fmt.println("success")
}
```
To run the sample:- go run example3.go

### example4.go:-(POOLING- Limit on the number of connections)

```go
package main

import (
           "database/sql"
           "fmt"
           "time"
          a "github.com/onlysumitg/GoQhttp"
       )

func ExecQuery(st *sql.Stmt) error {
        res, err := st.Query()
        if err != nil {
             fmt.Println(err)
        }
        cols, _ := res.Columns()

        fmt.Printf("%s    %s   %s     %s\n", cols[0], cols[1], cols[2], cols[3])
        defer res.Close()
        for res.Next() {
                    var t, x, m, n string
                    err = res.Scan(&t, &x, &m, &n)
                    fmt.Printf("%v  %v   %v     %v\n", t, x, m, n)
        }
        return nil
}

func main() {
	con := "HOSTNAME=host;PORT=number;DATABASE=name;UID=username;PWD=password"
        pool := a.Pconnect("PoolSize=5")

        ret := pool.Init(5, con)
        if ret != true {
                fmt.Println("Pool initializtion failed")
        }

        for i:=0; i<20; i++ {
                db1 := pool.Open(con, "SetConnMaxLifetime=10")
                if db1 != nil {
                        st1, err1 := db1.Prepare("select * from VMSAMPLE")
                        if err1 != nil {
                                fmt.Println("err1 : ", err1)
                        }else{
                                go func() {
                                        execquery(st1)
                                        db1.Close()
                                }()
                        }
                }
        }
        time.Sleep(30*time.Second)
        pool.Release()
}
```
To run the sample:- go run example4.go

For Running the Tests:
======================

1) Put your connection string in the main.go file in testdata folder

2) Now run go test command (use go test -v command for details) 

3) To run a particular test case (use "go test sample_test.go main.go", example "go test Arraystring_test.go main.go")
