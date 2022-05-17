package util

import (
	"crypto/md5"
	"encoding/hex"
)
//md5 the mima
func MD5(v string)string{
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

