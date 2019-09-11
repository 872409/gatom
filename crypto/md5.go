package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// Encrypt str
func MD5(str string) string {
	// hash := md5.Sum([]byte(str))
	// crypto := fmt.Sprintf("%x", hash)

	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
