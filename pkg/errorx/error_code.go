package errorx

import "strings"

type ErrorCode string

func (c ErrorCode) GetCodeMsg() (string, string) {
	str := string(c)
	index := strings.Index(str, ":")
	return str[:index], str[index+1:]
}
