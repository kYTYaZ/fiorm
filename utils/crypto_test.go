package utils_test

import (
	"math/rand"
	"strconv"
	"testing"

	"git.code.oa.com/fip-team/fiorm/utils"
)

func TestCrypto(t *testing.T) {
	password := "abc123456"
	//取新盐值
	salt := strconv.Itoa(int(rand.Int31n(2000000201)))
	//加密
	cryptoedPwd := utils.Crypto(password, salt)

	println(cryptoedPwd)
}
