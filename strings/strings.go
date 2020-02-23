package strings

import (
	"github.com/wide288/go-utils/converter"
	"strings"
)

// 判断子字符串是否存在于字符串中
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// 删除字符串两端空格符,空格符可以是连续多个
func TrimSpace(str string) string {
	return strings.Trim(str, " ")
}

// 删除字符串右侧自定义字符,删除的同一字符连续时数量不限
func TrimRight(str, cutset string) string {
	return strings.TrimRight(str, cutset)
}

func Split(str, sep string) []string {
	return strings.Split(str, sep)
}

func SplitInt(str, sep string) []int64 {
	strArr := strings.Split(str, sep)

	var intArr []int64
	for _, item := range strArr {
		intArr = append(intArr, converter.StrToInt64(item))
	}
	return intArr
}

func Join(strArr []string, sep string) string {
	return strings.Join(strArr, sep)
}