package main

// 在struct_table查看结构体
import (
	// mysql 驱动

	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB ss
var DB *sqlx.DB

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

// Schema 初始化数据库
type Schema struct {
	create string
	drop   string
}

var defaultSchema = Schema{
	create: `
		CREATE TABLE users (
		id int unsigned NOT NULL AUTO_INCREMENT,
		name varchar(255),
		added_at timestamp default now(),
		PRIMARY KEY (id)
		);
		CREATE TABLE userinfos (
			id int unsigned NOT NULL AUTO_INCREMENT,
			age int,
			user_id int,
			PRIMARY KEY (id)
		);
		CREATE TABLE classes (
			id int unsigned NOT NULL AUTO_INCREMENT,
			name varchar(256),
			PRIMARY KEY (id)
		);
		CREATE TABLE classes_users (
			id int unsigned NOT NULL AUTO_INCREMENT,
			user_id int,
			classes_id int,
			PRIMARY KEY (id)
		);
		CREATE TABLE books (
			id int unsigned NOT NULL AUTO_INCREMENT,
			name varchar(256),
			user_id int,
			PRIMARY KEY (id)
		);
		`,
	drop: `
		drop table users;
		drop table userinfos;
		drop table classes;
		drop table classes_users;
		drop table books;
		`,
}

// User 用户表
type User struct {
	ID      int
	Name    string
	AddedAt time.Time `db:"added_at"`
}

// Userinfo 用户详细信息，一对一用户表
type Userinfo struct {
	ID   int
	Age  int
	User User
}

// Class 班级表, 多对多
type Class struct {
	ID    int
	Users []*User
}

// Books 书本表, 用户一对多书本
type Books struct {
	ID   int
	User User
}

// MultiExec 批量执行
func MultiExec(e sqlx.Execer, query string) {
	stmts := strings.Split(query, ";\n")
	if len(strings.Trim(stmts[len(stmts)-1], " \n\t\r")) == 0 {
		stmts = stmts[:len(stmts)-1]
	}
	for _, s := range stmts {
		_, err := e.Exec(s)
		if err != nil {
			fmt.Println(err, s)
		}
	}
}

// RunWithSchema 创建表结构
func RunWithSchema(schema Schema, db *sqlx.DB, test func()) {

	// defer MultiExec(db, schema.drop)
	MultiExec(db, schema.drop)
	MultiExec(db, schema.create)
	test()
}

func loadDefaultFixture() {
	tx := DB.MustBegin()
	tx.MustExec(tx.Rebind("INSERT  users (id, name) VALUES (?, ?);"), "1", "wangmeimei")
	tx.MustExec(tx.Rebind("insert into users (id, name) values(?, ?);"), "2", "lisi")
	tx.MustExec(tx.Rebind("insert into users (id, name) values(?, ?);"), "3", "zhangsan")

	tx.MustExec(tx.Rebind("insert into userinfos (id, age, user_id) values(?, ?, ?);"), "1", "25", "1")
	tx.MustExec(tx.Rebind("insert into userinfos (id, age, user_id) values(?, ?, ?);"), "2", "18", "2")
	tx.MustExec(tx.Rebind("insert into userinfos (id, age, user_id) values(?, ?, ?);"), "3", "23", "3")

	tx.MustExec(tx.Rebind("insert into classes (id, name) values(?, ?);"), "1", "第三班")
	tx.MustExec(tx.Rebind("insert into classes (id, name) values(?, ?);"), "2", "第7班")
	tx.MustExec(tx.Rebind("insert into classes (id, name) values(?, ?);"), "3", "第88班")

	tx.MustExec(tx.Rebind("insert into classes_users (id, user_id, classes_id) values(?,?,?);"), "1", "1", "1")
	tx.MustExec(tx.Rebind("insert into classes_users (id, user_id, classes_id) values(?,?,?);"), "2", "2", "1")

	tx.MustExec(tx.Rebind("insert into books (id, name, user_id) values(?,?, ?);"), "1", "三体", "1")
	tx.MustExec(tx.Rebind("insert into books (id, name, user_id) values(?,?, ?);"), "2", "三年模拟", "1")
	tx.MustExec(tx.Rebind("insert into books (id, name, user_id) values(?,?, ?);"), "3", "三年模拟", "3")

	tx.Commit()
}

func test() {

}

func main() {
	Init("root:root1234@tcp(localhost:3306)/example?parseTime=true")
	defer DB.Close()

	RunWithSchema(defaultSchema, DB, test)
	loadDefaultFixture()
}
