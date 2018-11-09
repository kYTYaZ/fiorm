package example_test

import (
	"fmt"
	"testing"

	"git.code.oa.com/fip-team/fiorm"
	"git.code.oa.com/fip-team/fiorm/model"
)

func TestGetItemByID(t *testing.T) {
	var user model.User

	da := fiorm.DataAccept()
	da.GetItemByID(&user, 5)

	fmt.Println(user.Name)
}
