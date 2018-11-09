package example_test

import (
	"testing"

	"git.code.oa.com/fip-team/fiorm"

	"git.code.oa.com/fip-team/fiorm/model"
)

func TestUpdateOne(t *testing.T) {
	var user model.User
	da := fiorm.DataAccept()
	da.GetItemByID(&user, 5)

	user.Address = "udate1"

	da.UpdateItem(&user)
}
