package database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql" //加载mysql驱动
	"github.com/jinzhu/gorm"
	// "sfglxt/app/config"
    //_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func Init()  {
    var err error
    Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/sfglxt?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if Db.Error != nil {
        fmt.Printf("database error %v", Db.Error)
    }
}
// func Init() {
// 	fmt.Println("数据库持久化服务器启动:")
// 	Db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
// 	if err != nil {
// 		fmt.Println("open db connect error.", err)
// 		os.Exit(-1)
// 	}
// 	Db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
// 	Db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
// }