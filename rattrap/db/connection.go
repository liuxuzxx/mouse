package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"mouse/rattrap/config"
)

//
// 数据库信息的连接建立
//
var Db *gorm.DB

func init() {
	dbConfig := config.Conf.Database
	connectionUrl := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.DbName, dbConfig.Password)
	var err error
	Db, err = gorm.Open("postgres", connectionUrl)
	if err != nil {
		panic(err)
	}
	log.Printf("Init database connection success!")
}
