package db

import (
	"fmt"
	"testing"
)

type User struct {
	Id      string
	Name    string
	LoginId string `db:"login_id"`
}

func TestQuery(t *testing.T) {
	var u *User = new(User)
	err := Db.Get(u, "select id,name,login_id from sys_user where login_id='admin'")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	fmt.Println(u)
}

func TestExec(t *testing.T) {
	rs, err := Db.Exec("select * from sys_user")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	affected, err := rs.RowsAffected()
	fmt.Printf("exec success! rowsAffected -> %v", affected)
}
