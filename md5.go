package u

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// MD5 MD5加密
func MD5(data string) string {
	bytes := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", bytes)
}

// MD5Salt MD5加盐
func MD5Salt(data string, salt string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	hash.Write([]byte(salt))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

