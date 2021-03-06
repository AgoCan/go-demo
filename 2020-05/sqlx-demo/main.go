package main

// 在struct_table查看结构体
import (
	// mysql 驱动

	"database/sql"
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
	Name    string    `db:"name"`
	AddedAt time.Time `db:"added_at"`
}

// Userinfo 用户详细信息，一对一用户表
type Userinfo struct {
	ID     int
	Age    int
	UserID int `db:"user_id"`
}

// Class 班级表, 多对多
type Class struct {
	ID    int
	Users User `db:"users"`
	Name  string
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
func RunWithSchema(schema Schema, db *sqlx.DB, loadDefaultFixture func()) {

	// defer MultiExec(db, schema.drop)
	MultiExec(db, schema.drop)
	MultiExec(db, schema.create)
	loadDefaultFixture()
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
	tx.MustExec(tx.Rebind("insert into classes_users (id, user_id, classes_id) values(?,?,?);"), "3", "2", "2")

	tx.MustExec(tx.Rebind("insert into books (id, name, user_id) values(?,?, ?);"), "1", "三体", "1")
	tx.MustExec(tx.Rebind("insert into books (id, name, user_id) values(?,?, ?);"), "2", "三年模拟", "1")
	tx.MustExec(tx.Rebind("insert into books (id, name, user_id) values(?,?, ?);"), "3", "三年模拟", "3")

	tx.Commit()
}

func getOneRecord() {
	// 获取单条数据
	var user User
	err := DB.Get(&user, "select * FROM users LIMIT 1")
	if err != nil {
		panic(err)
	}

	rowsx, err := DB.Queryx("SELECT * FROM users ORDER BY id DESC LIMIT 1 ")
	if err != nil {
		panic(err)
	}
	rowsx.Next()

	err = rowsx.StructScan(&user)
	if err != nil {
		panic(err)
	}
	rowsx.Close()
	fmt.Println(user)

}

func multiRecord() {
	var users []User
	err := DB.Select(&users, "select * FROM users")
	if err != nil {
		panic(err)
	}
	// for _, v := range users {
	// 	fmt.Println(v.Name)
	// }

	var users02 []User
	db := DB.Unsafe()
	nstmt, err := db.PrepareNamed(`SELECT * FROM users WHERE name != :name`)
	if err != nil {
		panic(err)
	}
	err = nstmt.Select(&users02, map[string]interface{}{"name": "lisi"})
	if err != nil {
		panic(err)
	}
	for _, v := range users02 {
		fmt.Println(v.Name)
	}

}

func oneToOne() {
	// 一对一查询,也可以用与一对多查询
	userInfos := []struct {
		User
		Userinfo
	}{}
	var err error
	// 帮你自动查询两张表中所有相同的字段，然后进行等值连接。
	// err = DB.Select(
	// 	&userInfos,
	// 	`SELECT users.*, userinfos.* FROM
	// 	 users natural join userinfos`)
	// if err != nil {
	// 	panic(err)
	// }

	// err = DB.Select(
	// 	&userInfos,
	// 	`SELECT * FROM
	// 	users, userinfos
	// 	WHERE users.id = userinfos.user_id;`)
	// if err != nil {
	// 	panic(err)
	// }

	// err = DB.Select(
	// 	&userInfos,
	// 	`SELECT * FROM
	// 	users LEFT JOIN userinfos USING  (id);`)
	// if err != nil {
	// 	panic(err)
	// }

	// err = DB.Select(
	// 	&userInfos,
	// 	`SELECT * FROM
	// 	users INNER JOIN userinfos ON users.id = userinfos.user_id;`)
	// if err != nil {
	// 	panic(err)
	// }

	// 利用单表查询，并做组合
	var (
		users     []User
		userinfos []Userinfo
	)
	err = DB.Select(
		&users,
		`SELECT * FROM users`)
	if err != nil {
		panic(err)
	}
	err = DB.Select(
		&userinfos,
		`SELECT * FROM userinfos`)
	if err != nil {
		panic(err)
	}

	for _, v1 := range users {
		for _, v2 := range userinfos {
			if v2.UserID == v1.ID {
				userInfos = append(userInfos, struct {
					User
					Userinfo
				}{User: v1, Userinfo: v2})
			}
		}
	}

	fmt.Println(userInfos)
}

func manyToMany() {
	// 多对多，三表查询
	// var classes []Class // 结构体是什么类型，读出来就是什么数据
	// var err error
	// err = DB.Select(
	// 	&classes,
	// 	`SELECT classes.id id, users.id "users.id", users.name  "users.name", classes.name  FROM users
	// 	INNER JOIN classes
	// 	INNER JOIN classes_users
	// 	ON classes.id = classes_users.classes_id and users.id = classes_users.user_id ;`)
	// //
	// if err != nil {
	// 	panic(err)
	// }

	type Class02 struct {
		ID    int
		Users []*User
		Name  string
	}
	var classList []*Class02

	var classesMap = make(map[int][]*User)
	rows, _ := DB.Queryx(`SELECT classes.id id, users.id "users.id", users.name  "users.name", classes.name  FROM users 
	INNER JOIN classes 
	INNER JOIN classes_users 
	ON classes.id = classes_users.classes_id and users.id = classes_users.user_id ;`)
	for rows.Next() {
		var (
			ClassID   int
			UserID    int
			UserName  string
			ClassName string
		)
		rows.Scan(&ClassID, &UserID, &UserName, &ClassName)
		if users, ok := classesMap[ClassID]; ok {
			classesMap[ClassID] = append(users, &User{ID: UserID, Name: UserName})

		} else {
			classesMap[ClassID] = []*User{&User{ID: UserID, Name: UserName}}
			classList = append(classList, &Class02{ID: ClassID, Name: ClassName})
		}

	}

	for _, c := range classList {
		if users, ok := classesMap[c.ID]; ok {
			c.Users = users
		}
	}

	for _, c := range classList {
		fmt.Println(c)
	}
}

func useNullString() {
	// 如果注释掉下面的代码，使用最上面的结构体，那么就会报错 panic: sql: Scan error on column index 1, name "name": converting NULL to string is unsupported
	type User struct {
		ID      int
		Name    sql.NullString
		AddedAt time.Time `db:"added_at"`
	}
	var user User
	var err error
	err = DB.Get(&user, "select * FROM users LIMIT 1")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	// UPDATE dockerfile SET dockerfile = ? where level_id = ?
	_, err = DB.Exec("UPDATE users SET name = ? where id = 1", sql.NullString{String: "", Valid: false})
	fmt.Println(err)
	err = DB.Get(&user, "select * FROM users LIMIT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

func useIn() {
	var users []*User
	sqlStr := `select 
		*
	from 
		users
	where 
		id in (?); `
	// 其中 query 是字符串。根据数组的长度拆开了问好
	query, args, err := sqlx.In(sqlStr, []int{1, 2})
	fmt.Println(query, args)
	if err != nil {
		panic(err)
	}
	rows, err := DB.Queryx(query, args...)
	for rows.Next() {
		var user User
		if rows.StructScan(&user) != nil {
			break
		}
		users = append(users, &user)
	}
	fmt.Println(users)

}

func main() {
	Init("root:root1234@tcp(localhost:3306)/example?parseTime=true")
	defer DB.Close()

	RunWithSchema(defaultSchema, DB, loadDefaultFixture)
	useIn()
}
