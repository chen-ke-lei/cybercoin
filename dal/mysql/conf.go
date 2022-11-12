package mysql

import (
	"cybercoin/dal/const_dal"
	"cybercoin/dal/data_object"
	"cybercoin/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySql *gorm.DB

func init() {
	mysqlConf := Conf{}
	util.ReadTransfer(const_dal.MYSQL_CONF, &mysqlConf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DataBase,
		mysqlConf.Timeout)
	MySql, _ = gorm.Open(mysql.Open(dsn))
	MySql.AutoMigrate(&data_object.CoinHisPrice{})

}

type Conf struct {
	Username string //账号
	Password string //密码
	Host     string //数据库地址，可以是Ip或者域名
	Port     int    //数据库端口
	DataBase string //数据库名
	Timeout  string //连接超时，10秒
	MaxIdle  int
}
