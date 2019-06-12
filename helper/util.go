//工具包
package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
)

//封装一个分页结构体
type PageHelper struct {
	TotalCount       int64 //总页数
	PageSize         int64 //每页记录数
	PerSize          int64 //总页码
	CurrentPageIndex int64 //当前页数
	PrePageIndex     int64 //上一页
	NextPageIndex    int64 //下一页
}

//分页结构体初始化
func PageHelperInit(totalCount int64, perSize int64, currentPageIndex int64) *PageHelper {
	pg := new(PageHelper)
	pg.TotalCount = totalCount
	pg.PerSize = perSize

	if currentPageIndex <= 1 {
		pg.PrePageIndex = 1
	}
	if currentPageIndex >= pg.PageSize {
		pg.NextPageIndex = totalCount
	}
	return pg
}

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

//返回一个md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//返回16位md5加密后的字符串
func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}

func GetSpiltLastStr(src string) string {
	slices := strings.Split(src, "/")
	fmt.Println(slices)
	lastStr := slices[len(slices)-1]
	return lastStr
}
