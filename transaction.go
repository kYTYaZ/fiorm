package fiorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// TxDB 事务
type TxDB struct {
	db    *gorm.DB // 局部
	Error error
}

// BeginTx 开始事务
func (t *FiDB) BeginTx() *TxDB {
	var tx = new(TxDB)
	tx.db = db.Begin()

	return tx
}

// EndTx 结束事务
func (t *TxDB) EndTx() {
	if t.db.Error != nil || t.Error != nil {
		t.db.Rollback()
		fmt.Println(t.Error)
		panic(t.db.Error)
	}

	// TODO t.db.GetErrors ??
	t.db.Commit()
}

// InsertItem 插入一条或多条数据
func (t *TxDB) InsertItem(value interface{}) {
	d := t.db.Create(value)

	if d.Error != nil {
		panic(d.Error)
	}
}
