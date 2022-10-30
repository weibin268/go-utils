package db

import (
	"fmt"
	"log"
	"sort"
	"testing"
)

type User struct {
	Id      string
	Name    string
	LoginId string `db:"login_id"`
}

func TestGet(t *testing.T) {
	InitDb()
	defer CloseDb()
	var u *User = new(User)
	err := Db.Get(u, "select id,name,login_id from sys_user where login_id='admin'")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	fmt.Println(u)
}

func TestQueryMaps(t *testing.T) {
	InitDb()
	defer CloseDb()
	rows, err := QueryMaps("select * from sys_user")
	if err != nil {
		log.Fatal(err)
	}
	var colums []string
	for _, r := range rows {
		if colums == nil {
			for k, _ := range r {
				colums = append(colums, k)
			}
			sort.Strings(colums)
			for _, c := range colums {
				fmt.Printf("%v|", c)
			}
			fmt.Printf("\n")
		}
		for _, k := range colums {
			s := r[k]
			fmt.Printf("%v|", s)
		}
		fmt.Printf("\n")
	}
}

func TestExec(t *testing.T) {
	InitDb()
	defer CloseDb()
	rs, err := Db.Exec("select * from sys_user")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	affected, err := rs.RowsAffected()
	fmt.Printf("exec success! rowsAffected -> %v", affected)
}
