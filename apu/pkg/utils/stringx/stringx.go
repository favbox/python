package stringx

import (
	"regexp"
	"strings"
	"unicode"
)

// Trim 删除左右两侧及零宽空格。
func Trim(str string) string {
	// Remove zero-width spaces
	str = strings.Replace(str, "\u200B", "", -1)
	str = strings.Replace(str, "\u00a0", "", -1)
	return strings.TrimSpace(str)
}

// RemoveAllSpace 删除字符串中的所有空格。
func RemoveAllSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

// ReduceEmptyLines 缩减空白行。
func ReduceEmptyLines(input string) string {
	// 使用正则表达式匹配两个及以上连续的空行
	pattern := regexp.MustCompile(`\n{2,}`)
	// 使用 ReplaceAllString 方法将匹配到的连续空行替换为一个换行符
	output := pattern.ReplaceAllString(input, "\n")
	// 返回处理后的结果
	return output
}

// Replace 使用正则替换字符串。
func Replace(str string, re *regexp.Regexp, repl ...string) string {
	r := ""
	if len(repl) > 0 {
		r = repl[0]
	}
	str = Trim(re.ReplaceAllString(str, r))
	return str
}