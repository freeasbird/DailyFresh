//工具包
package helper

import (
	"fmt"
	"github.com/satori/go.uuid"
)

//返回一个uuid字符串
func getUUID(version int) string {
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("UUIDv4: %s\n", uid)
	}
	return uid.String()
}
