package util

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

type AnsInfo struct {
	Author   string
	Title    string
	Content  string
	PostTime time.Time
	Counter  int
	Id       string
}

func ContentSha1(s string) string {
	// 创建 SHA-1 哈希对象
	hash := sha1.New()

	// 写入要计算哈希的数据
	hash.Write([]byte(s))

	// 计算哈希值
	hashValue := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashValue)

	return hashString
}
