package db

import (
	"fmt"
	"log"
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
	maps, err := QueryMaps("select * from sys_user")
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range maps {
		for k, v := range m {
			fmt.Printf("%v:%v|", k, v)
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
