package main

import (
	"fmt"
	"time"
)

func formattedStringToLDAPTimestamp(formattedTime string) string {
	// 解析格式化的日期时间字符串为 time.Time 对象
	parsedTime, err := time.Parse("2006/1/2 15:04:05", formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return ""
	}

	// 1601-01-01 00:00:00 UTC 对应的纳秒数
	const filetimeEpoch = 116444736000000000

	// 获取时间对象对应的纳秒数并计算与FileTime的起始值的差
	totalNanoseconds := parsedTime.UnixNano()
	ldapTimestamp := totalNanoseconds/100 + filetimeEpoch

	// 转换为字符串并返回18-digit LDAP timestamp格式
	return fmt.Sprintf("%018d", ldapTimestamp)
}

func main() {
	// 给定的格式化日期时间字符串
	formattedTimeString := "2024/1/15 00:00:00"

	// 将格式化的日期时间字符串转换为18-digit LDAP timestamp
	ldapTimestamp := formattedStringToLDAPTimestamp(formattedTimeString)

	// 输出转换后的18-digit LDAP timestamp值
	fmt.Println("18-digit LDAP Timestamp:", ldapTimestamp)
}
