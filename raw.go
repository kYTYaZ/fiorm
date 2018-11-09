package fiorm

// 原生数据库支持
import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Raw 原生SQL
func Raw(sql string, values ...interface{}) (rows *sql.Rows) {
	d := db.Raw(sql, values...)
	var err error
	rows, err = d.Rows()

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return rows
}
