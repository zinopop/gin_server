package model

import (
	"fmt"
	"gin_server/lib"
)

type UserTest struct {
	Id   int64
	Name string
}

func (UserTest) NewUser(Name string) (int64, error) {
	engine := lib.GetConn()
	return engine.Insert(&UserTest{
		Name: "测试",
	})
}

func (UserTest) GetUserById(Id int64) (UserTest, error) {
	var userTest UserTest
	engine := lib.GetConn()
	_, err := engine.Where("id = ?", Id).Get(&userTest)
	fmt.Println(userTest)
	if err != nil {
		return userTest, err
	}
	return userTest, nil
}
