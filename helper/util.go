//工具包
package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
)

//返回一个uuid字符串
func GetUUID(version int) string {
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("UUIDv4: %s\n", uid)
	}
	return uid.String()
}

//返回一个sha256加密后的字符串
func GetSha256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
