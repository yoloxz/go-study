package utils

import "unicode"

// 判断字符串是否只包含中文、字母、数字、下划线
func IsLetterOrDigit(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '_' {
			return false
		}
	}
	return true
}
