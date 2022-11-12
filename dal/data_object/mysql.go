package data_object

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySql *gorm.DB

func init() {
	username := "root"          //账号
	password := "root"          //密码
	host := "127.0.0.1"         //数据库地址，可以是Ip或者域名
	port := 3306                //数据库端口
	Dbname := "quant_cybercoin" //数据库名
	timeout := "10s"            //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	MySql, _ = gorm.Open(mysql.Open(dsn))
	MySql.AutoMigrate(&CoinHisPrice{})
}
