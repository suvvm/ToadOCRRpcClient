package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func GetBasicToken(appSecret, verifyStr string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(appSecret + verifyStr))
	md5Token := hex.EncodeToString(hasher.Sum(nil))
	return md5Token, nil
}

func PixelHashStr(pxs []byte) string {
	var resp string
	for i := 0; i < len(pxs); i++ {
		pixelVal := int(pxs[i])
		resp += strconv.Itoa(pixelVal)
	}
	return resp
}
