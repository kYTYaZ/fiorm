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
	db    *gorm.DB // 全局DB
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
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, p, dbname)
	if source == "" {
		panic("错误的连接字符串")
	}

	var err error
	db, err = gorm.Open(dialect, source)
	if err != nil {
		panic("数据库没有初始化")
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

// DataAccept 获取一个数据库连接
func DataAccept() *FiDB {
	var da = new(FiDB)
	da.db = db
	return da
}

// GetItemByID 根据主键ID获取数据
func (t *FiDB) GetItemByID(tEntity interface{}, id int64) {
	d := db.First(tEntity, id)

	if d.Error != nil && d.Error != gorm.ErrRecordNotFound {
		panic(d.Error)
	}
}

// GetItemWhereFirst 根据条件查询一条数据
func (t *FiDB) GetItemWhereFirst(tEntity interface{}, query *Query) {
	if query.db.Error != nil {
		panic(query.db.Error)
	}

	d := query.db.First(tEntity)

	if d.Error != nil && d.Error != gorm.ErrRecordNotFound {
		panic(d.Error)
	}
}

// GetItemWhere 根据条件查询多条数据
func (t *FiDB) GetItemWhere(tEntity interface{}, query *Query) {
	if query.db.Error != nil {
		panic(query.db.Error)
	}

	d := query.db.Find(tEntity)

	if d.Error != nil && d.Error != gorm.ErrRecordNotFound {
		panic(d.Error)
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
	d := t.db.Create(value)

	if d.Error != nil {
		panic(d.Error)
	}
}
