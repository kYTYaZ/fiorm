package fiorm

import (
	"fmt"
	"log"
	"os"
	"strconv"

	// 调用mysql init方法初始化
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 返回一个连接池的实例
var db *gorm.DB

// FiDB fiorm处理类
type FiDB struct {
	db    *gorm.DB // 局部
	Error error
}

// InitDB 初始化
//
// dialect -别名,如mysql,oracle
// InitDB("mysql", "GO_TESTDB", "localhost", "root", "password", 3306)
//
// The returned DB is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a DB.
func InitDB(dialect string, dbname string, host string, user string, password string, port int) {
	p := strconv.Itoa(port)
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=true", user, password, host, p, dbname)
	if source == "" {
		panic("错误的连接字符串")
	}

	var err error
	db, err = gorm.Open(dialect, source)
	if err != nil {
		panic("数据库没有初始化")
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

// DataAccept 获取一个数据库连接
func DataAccept() *FiDB {
	var da = new(FiDB)
	return da
}

// GetItemByID 根据主键ID获取数据
func (t *FiDB) GetItemByID(tEntity interface{}, id int64) {
	t.db = db.First(tEntity, id)

	if t.db.Error != nil && t.db.Error != gorm.ErrRecordNotFound {
		panic(t.db.Error)
	}
}

// GetItemWhereFirst 根据条件查询一条数据
func (t *FiDB) GetItemWhereFirst(tEntity interface{}, query *Query) {
	if query.db.Error != nil {
		panic(query.db.Error)
	}

	t.db = query.db.First(tEntity)

	if t.db.Error != nil && t.db.Error != gorm.ErrRecordNotFound {
		panic(t.db.Error)
	}
}

// GetItemWhere 根据条件查询多条数据
func (t *FiDB) GetItemWhere(tEntity interface{}, query *Query) {
	if query.db.Error != nil {
		panic(query.db.Error)
	}

	t.db = query.db.Find(tEntity)

	if t.db.Error != nil && t.db.Error != gorm.ErrRecordNotFound {
		panic(t.db.Error)
	}
}

// Count 返回总行数
func (t *FiDB) Count(value interface{}) {
	if t.db.Error != nil && t.db.Error != gorm.ErrRecordNotFound {
		panic(t.db.Error)
	}

	t.db.Count(value)
}

// InsertItem 插入一条或多条数据
func (t *FiDB) InsertItem(value interface{}) {
	t.db = t.db.Create(value)

	if t.db.Error != nil {
		panic(t.db.Error)
	}
}

// DeleteItem 删除一条或多条数据
func (t *FiDB) DeleteItem(value interface{}) {
	if t.db.Error == gorm.ErrRecordNotFound {
		return
	}
	// 防止删除所有数据
	if t.db.Error != nil {
		panic(t.db.Error)
	}
	if t.Error != nil {
		panic(t.Error)
	}

	// 必须用t.db而不是全局db,全局db删除的情况下，如果value没有值，将导致删除所有的数据！
	t.db.Delete(value)
}

// UpdateItem 更新一行或多行数据
func (t *FiDB) UpdateItem(value interface{}) {
	if t.db.Error == gorm.ErrRecordNotFound {
		return
	}
	// 防止更新所有数据
	if t.db.Error != nil {
		panic(t.db.Error)
	}
	if t.Error != nil {
		panic(t.Error)
	}

	// 必须用t.db而不是全局db,全局db更新的情况下，如果value没有值，将导致更新所有的数据！
	t.db.UpdateColumns(value)
}

// CreateTable 创建表 DDL操作
func CreateTable(value interface{}) {
	db.CreateTable(value)
}
