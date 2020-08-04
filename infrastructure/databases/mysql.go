package databases

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Mysql dbをmysqlにする場合はこれを使う
type Mysql struct {
}

//NewMysql New Mysql
func NewMysql() *Mysql {
	return &Mysql{}
}

// GormConnect MySQL wrapper に接続
func (mysql *Mysql) Open() *gorm.DB {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	option := "charset=utf8&parseTime=True&loc=Asia%2FTokyo"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", user, password, host, name, option)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
