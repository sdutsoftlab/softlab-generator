package utils

import (
	"strings"
)

// 小写转大写  空格换-
func Convert(str string) string {
	str = strings.ToLower(str)
	ss := strings.SplitN(str, " ", -1)
	return strings.Join(ss, "-")
}
