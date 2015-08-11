package common

import (
	"crypto/md5"
	"encoding/hex"
)

func MakeMD5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}
