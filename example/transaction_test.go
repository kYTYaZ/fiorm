package example_test

import (
	"database/sql"
	"errors"
	"sync"
	"testing"
	"time"

	"git.code.oa.com/fip-team/fiorm"
	"git.code.oa.com/fip-team/fiorm/model"
	"github.com/go-sql-driver/mysql"
)

//事务测试
func TestSimpleTx(t *testing.T) {
	var user model.User
	user.Address = "深圳南山TX"
	user.CreatedAt = time.Now()
	user.DeptId = sql.NullInt64{Int64: 0, Valid: false}
	user.Email = ""
	user.Name = "wins"
	user.NullAge = sql.NullInt64{Int64: 21, Valid: true}
	user.NullString = ""
	user.Birthday = mysql.NullTime{Time: time.Now(), Valid: false}

	var user2 model.User
	user2.Address = "深圳南山TX2"
	user2.CreatedAt = time.Now()
	user2.DeptId = sql.NullInt64{Int64: 22, Valid: true}
	user2.Email = ""
	user2.Name = "winstx"
	user2.NullAge = sql.NullInt64{Int64: 222, Valid: true}
	user2.NullString = ""
	user2.Birthday = mysql.NullTime{Time: time.Now(), Valid: true}

	var user3 model.User
	user3.Address = "深圳南山TX3"
	user3.CreatedAt = time.Now()
	user3.DeptId = sql.NullInt64{Int64: 0, Valid: false}
	user3.Email = ""
	user3.Name = "wins"
	user3.NullAge = sql.NullInt64{Int64: 21, Valid: true}
	user3.NullString = ""
	user3.Birthday = mysql.NullTime{Time: time.Now(), Valid: false}

	da := fiorm.DataAccept()
	tx := da.BeginTx()

	tx.InsertItem(&user3)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go insert(tx, &user, &wg, 0)
	}

	wg.Add(1)
	go insert(tx, &user2, &wg, 0)

	wg.Wait()
	tx.EndTx()
}

func insert(tx *fiorm.TxDB, value interface{}, wg *sync.WaitGroup, idx int) {
	tx.InsertItem(value)
	wg.Done()

	if idx == 1 {
		tx.Error = errors.New("测试错误")
	}
}
