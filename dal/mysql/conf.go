package mysql

import (
	"context"
	"cybercoin/dal/const_data"
	"cybercoin/dal/data_object"
	"cybercoin/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var MySql *gorm.DB
var once sync.Once
var MysqlConf Conf

func NewClient(ctx context.Context) *gorm.DB {
	once.Do(func() {
		MysqlConf := Conf{}
		util.ReadTransfer(const_data.MYSQL_CONF, &MysqlConf)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
			MysqlConf.Username,
			MysqlConf.Password,
			MysqlConf.Host,
			MysqlConf.Port,
			MysqlConf.DataBase,
			MysqlConf.Timeout)
		MySql, _ = gorm.Open(mysql.Open(dsn))
		MySql.AutoMigrate(&data_object.CoinHisPrice{})
	})
	return MySql
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
