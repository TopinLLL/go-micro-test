package conf

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func InitDB()(err error){
	username:=viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",username,password,host,port,database,charset)
	DB,err=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println("数据库连接错误",err)
	}
	return err
}
