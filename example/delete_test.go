package example_test

import (
	"testing"

	"git.code.oa.com/fip-team/fiorm"

	"git.code.oa.com/fip-team/fiorm/model"
)

func TestDeleteOne(t *testing.T) {
	var user model.User
	da := fiorm.DataAccept()
	da.GetItemByID(&user, 6)

	da.DeleteItem(&user)
}
