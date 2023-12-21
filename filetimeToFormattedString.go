package main

import (
	"fmt"
	"time"
)

func filetimeToFormattedString(filetime uint64) string {
	// 1601-01-01 00:00:00 UTC 对应的 FileTime 值（以纳秒为单位）
	const filetimeEpoch = 116444736000000000

	// 将 FileTime 转换为纳秒并减去起始的差值
	ns := int64(filetime) - filetimeEpoch

	// 将纳秒转换为秒
	secs := ns / 10000000 // 100纳秒 = 0.0000001秒，所以10^7纳秒 = 1秒

	// 使用秒数创建 time.Time 对象
	timeObj := time.Unix(secs, 0)

	// 将 time.Time 对象格式化为所需的字符串格式
	return timeObj.Format("2006/1/2 15:04:05")
}

func main() {
	// 给定的 FileTime 值，根据您提供的参考链接，这个值可能需要相应的转换
	filetimeValue := uint64(133497216000000000)

	// 将 FileTime 转换为格式化的字符串
	formattedString := filetimeToFormattedString(filetimeValue)

	// 输出转换后的日期时间
	fmt.Println("Converted Time:", formattedString)
}
