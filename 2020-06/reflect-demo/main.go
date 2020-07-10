package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type lowBase struct {
	ID         int `db:"low_id"`
	CreateTime int
}
type base struct {
	ID         float32 `db:"id"`
	CreateTime int
}

// User sd
type User struct {
	Base      base
	UserID    int64     `db:"user_id"`
	Username  string    `db:"username"  json:"username"`
	Password  string    `db:"password"  json:"password"`
	Email     string    `db:"email"     json:"email"`
	Telephone int       `db:"telephone" json:"telephone"`
	Avatar    string    `db:"avatar"    json:"avatar"`
	NickName  string    `db:"nick_name" json:"nick_name"`
	T         time.Time `db:"t" json:"t"`
}

func judgeType(v string) string {
	index := strings.Contains(v, "int")
	f := strings.Contains(v, "float")
	fmt.Println(index, f, v)
	var str string
	if index {
		str = "%v"
	} else {
		if f {
			str = "%v"
		} else {
			str = "'%v'"
		}

	}
	return str

}

// ReflectTag 构建  tag=值的字符串， 只能两层结构体
func reflectTag(s interface{}, arg ...string) (resArray []string, err error) {
	var res string

	st := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		// 直接不添加到切片中，这样的目的
		if arg != nil {
			flag := false
			for k := 0; k < len(arg); k++ {
				if field.Name == arg[k] {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
		}
		if field.Name == "baseModel" {
			continue
		}

		if !v.Field(i).IsZero() {
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					if structField.Field(j).Tag.Get("db") == "" {
						continue
					}
					str := "%s=" + judgeType(fmt.Sprintf("%v", v.Field(i).Field(j).Type()))
					res = fmt.Sprintf(str, structField.Field(j).Tag.Get("db"), v.Field(i).Field(j))
					resArray = append(resArray, res)
				}
				continue
			} else {
				str := "%s=" + judgeType(fmt.Sprintf("%v", v.Field(i).Type()))
				res = fmt.Sprintf(str, field.Tag.Get("db"), v.Field(i))
				resArray = append(resArray, res)
			}

		}
	}
	if len(resArray) == 0 {
		return resArray, errors.New("resArray is empty")
	}
	return resArray, nil
}
func main() {
	s := User{
		UserID:   11,
		Username: "123",
		Base:     base{ID: 1},
		T:        time.Now(),
	}
	res, err := reflectTag(s)
	sqlStr := strings.Join(res, ",")
	fmt.Println(sqlStr, err)
}
