package clang

import (
	"errors"
	"regexp"
	"strings"
)

//
//
// StringIgnoreEquals 忽略大小写比较
//
//
func StringIgnoreEquals(src string, dest string) bool {
	src = strings.ToLower(src)
	dest = strings.ToLower(dest)
	return src == dest
}

//
//
// StringSplitByRegexp 正则切割
//
//
func StringSplitByRegexp(src string, reg string) ([]string, error) {
	pattern := regexp.MustCompile(reg)
	indexes := pattern.FindAllStringIndex(src, -1)
	if len(indexes) == 0 {
		return nil, errors.New("no data")
	}
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = src[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = src[laststart:len(src)]
	return nil, result
}

//
//
// StringFirstLetterLower 首字母小写
//
//
func StringFirstLetterLower(src string) string {
	return strings.ToLower(src[:1]) + src[1:]
}

//
//
// StringFirstLetterUpper 首字母大写
//
//
func StringFirstLetterUpper(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

//
//
// StringTrim 取首尾空格
//
//
func StringTrim(src string) string {
	return strings.TrimSpace(src)
}

//
//
// StringBetween 取2个字符串中间的字符
//
//
func StringBetween(src string, start string, end string) (string, error) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		src = src[sI+len(start):]
		eI := strings.Index(src, end)
		if eI > 0 {
			return src[:eI], nil
		}
		return "", errors.New("end not found")
	}
	return "", errors.New("start not found")
}

//
//
// StringStartWith 是否是指定字符串开始
//
//
func StringStartWith(src string, s string) bool {
	return strings.Index(src, s) == 0
}

//
//
// StringAfter 获取指定字符之后的
//
//
func StringAfter(src string, start string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		return src[sI+len(start):], true
	}
	return "", false
}

//
//
// StringBefore 获取指定字符串之前的
//
//
func StringBefore(src string, start string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 1 {
		return src[:sI], true
	}
	return "", false
}

//
//
// StringMatch 获取指定正则中的group
//
//
func StringMatch(src string, p string, group int) (string, error) {
	pattern, err := regexp.Compile(p)
	if err != nil {
		return "", err
	}
	r := pattern.FindStringSubmatch(src)
	if len(r) > group {
		return r[group], nil
	}
	return "", errors.New("匹配失败")
}

//
//
// StringReplaceByRegex 正则替换
//
//
func StringReplaceByRegex(src string, p string, r string) string {
	rp := regexp.MustCompile(p)
	return rp.ReplaceAllString(src, r)
}

//
//
// StringLeftPad 左补齐
//
//
func StringLeftPad(src string, length int, pad string) string {
	srcLen := len(src)
	offset := length - srcLen
	for i := 0; i < offset; i++ {
		src = pad + src
	}
	return src
}
