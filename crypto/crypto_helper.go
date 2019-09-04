package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func StrMD5(str string) string {
	// hash := md5.Sum([]byte(str))
	// crypto := fmt.Sprintf("%x", hash)

	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

func StrSHA256(str string) string {
	ctx := sha256.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
