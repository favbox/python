package stringx

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/exp/constraints"
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
func ReduceEmptyLines(input string, n ...int) string {
	num := 3
	if len(n) > 0 {
		num = n[0]
	}
	if num < 1 {
		num = 1
	}

	// 使用正则表达式匹配3个及以上连续的空行
	pattern := regexp.MustCompile(fmt.Sprintf(`\n{%d,}`, num))
	// 使用 ReplaceAllString 方法将匹配到的连续空行替换为一个换行符
	output := pattern.ReplaceAllString(input, strings.Repeat("\n", num-1))
	// 返回处理后的结果
	return strings.TrimSpace(output)
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

type Parseable interface {
	// NOTE: I didn't check that fmt.Sscanf can accept all these,
	// but it seems like it probably should...
	string | bool | constraints.Integer | constraints.Float | constraints.Complex
}

func Parse[T Parseable](str string) (T, error) {
	var result T
	_, err := fmt.Sscanf(str, "%v", &result)
	return result, err
}

func MustNumber[T constraints.Float | constraints.Integer](s string) T {
	n, err := Parse[T](s)
	if err != nil {
		return 0
	}
	return n
}

// HasChinese 判断文本中是否有汉字。
func HasChinese(text string) bool {
	for _, r := range text {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// Cut 裁剪字符串
func Cut(s string, length int) string {
	var count int
	var result []rune
	for _, r := range s {
		if count >= length {
			break
		}
		if unicode.Is(unicode.Han, r) {
			count += 2
		} else {
			count++
		}
		result = append(result, r)
	}
	return string(result)
}
