package tools

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Tools 是内部单例实体
type Tools struct {
}

var (
	// Tool 为单例实体对外的名称
	Tool = New()
	once sync.Once
)

// New 返回单例实例
func New() (t *Tools) {
	once.Do(func() { //只执行一次
		t = &Tools{}
	})
	return t
}

// ParseInt string转换int
func (t *Tools) ParseInt(b string, defInt int) int {
	id, err := strconv.Atoi(b)
	if err != nil {
		return defInt
	}
	return id
}

// ParseString int转换string
func (t *Tools) ParseString(b int) string {
	id := strconv.Itoa(b)
	return id
}

// ParseFlostToString 转换浮点数为string
func (t *Tools) ParseFlostToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 5, 64)
}

// StructToString 结构体转成json字符串
func (t *Tools) StructToString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// StructToMap 结构体转换成map对象
func (t *Tools) StructToMap(obj interface{}) map[string]interface{} {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < k.NumField(); i++ {
		data[strings.ToLower(k.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

// GetRandomString 生成随机字符串
func (t *Tools) GetRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()+[]{}/<>;:=.,?"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// SubString 字符串截取
func (t *Tools) SubString(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	runeStr := []rune(str)
	lenStr := len(runeStr)

	if start < 0 {
		start = lenStr + start
	}
	if start > lenStr {
		start = lenStr
	}
	end := start + length
	if end > lenStr {
		end = lenStr
	}
	if length < 0 {
		end = lenStr + length
	}
	if start > end {
		start, end = end, start
	}
	return string(runeStr[start:end])
}

// Base64Decode base64 解码
func (t *Tools) Base64Decode(str string) string {
	s, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(s)
}

// TimeFormat 格式化时间
func (t *Tools) TimeFormat(time *time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
