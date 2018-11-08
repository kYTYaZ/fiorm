package model

import (
	"database/sql"
	"time"
)

// Department 部门表
type Department struct {
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	NullFloat sql.NullFloat64
	DeptName  string
	Tel       string `gorm:"type:varchar(100);not null"`
	Address   string `gorm:"size:255"` // 给address字段创建名为addr的索引
}
