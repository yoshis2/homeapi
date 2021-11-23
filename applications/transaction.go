package applications

import "gorm.io/gorm"

//Transact DBのトランザクション管理
func Transaction(db *gorm.DB, txFunc func(*gorm.DB) error) (err error) {
	tx := db.Begin()
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

	err = tx.Error
	if err != nil {
		return err
	}

	err = txFunc(tx)
	return err
}
