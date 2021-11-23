package databases

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Mysql dbをmysqlにする場合はこれを使う
type Mysql struct {
}

//NewMysql New Mysql
func NewMysql() *Mysql {
	return &Mysql{}
}

// GormConnect MySQL wrapper に接続
func (mysqls *Mysql) Open() *gorm.DB {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	option := "charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", user, password, host, name, option)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
