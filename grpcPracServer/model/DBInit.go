package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func MysqlInit(){
	sql,err :=gorm.Open("mysql","root:@tcp(127.0.0.1:3306)/ongorm?charset=utf8&parseTime=true")
	if err!=nil {
		fmt.Println(err.Error())
	}
	DB=sql
}


