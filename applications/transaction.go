package applications

import (
	"log"

	"gorm.io/gorm"
)

// Transact DBのトランザクション管理
func Transaction(db *gorm.DB, txFunc func(*gorm.DB) error) (err error) {
	log.Printf("トランザクション１ %v です", db)
	tx := db.Begin()
	log.Println("トランザクション2")
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	log.Println("トランザクション10")

	if err := tx.Error; err != nil {
		log.Printf("エラートランザクション : %v", err)
		return err
	}

	err = txFunc(tx)
	return err
}
