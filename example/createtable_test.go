package example_test

import (
	"testing"

	"git.code.oa.com/fip-team/fiorm"
	"git.code.oa.com/fip-team/fiorm/model"
)

// 创建表

func TestCreateTable(t *testing.T) {
	var user model.User
	var dept model.Department

	fiorm.CreateTable(&user)
	fiorm.CreateTable(&dept)

}
