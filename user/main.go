package main

import (
	"fmt"
	"user/conf"
	"user/model"

	"github.com/spf13/viper"
)

func main(){
	InitConfig()
	err:=conf.InitDB()
	if err != nil {
		println(err)
	}
	err = conf.DB.AutoMigrate(&model.User{})
	if err != nil {
		println("数据库建表失败")
	}
}


func InitConfig(){
	viper.SetConfigName("application")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}
