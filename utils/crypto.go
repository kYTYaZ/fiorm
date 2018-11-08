package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

//比MD5可靠很多很多的盐值加密算法
//加密速度还可以
//不怕撞库攻击
//不可逆加密
//在不知道盐值的情况下几乎不能进行暴力破解
//即使盐值一起泄露，如果不知道加密次数，暴力破解的成本极高

// Crypto 带盐值加密
func Crypto(password, salt string) string {
	prf := hmac.New(sha256.New, []byte(password))
	hashLen := prf.Size()
	numBlocks := (32 + hashLen - 1) / hashLen
	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		prf.Reset()
		prf.Write([]byte(salt))
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)
		for n := 2; n <= 10000; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			// fmt.Println(T)
			for x := range U {
				T[x] ^= U[x]
			}
			// fmt.Println(T)
		}
	}
	return fmt.Sprintf("%x", dk[:32])
}
