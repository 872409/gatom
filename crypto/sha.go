package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(str string) string {
	ctx := sha256.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
