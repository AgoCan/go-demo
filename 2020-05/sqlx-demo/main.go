package main

import (
	// mysql 驱动
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// DB ss
	DB *sqlx.DB
)

// Init ss
func Init(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

var (
	userTable = `insert into t1 values(
		1,
		'heihei'
	)
	`
)

func main() {
	Init("root:root1234@tcp(localhost:3306)/example?parseTime=true")
	defer DB.Close()

	res, err := DB.Exec(userTable)
	fmt.Println(res, err)
	fmt.Println(res.LastInsertId())
}
