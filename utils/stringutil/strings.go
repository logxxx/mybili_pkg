package stringutil

import "strings"

func TransHtml(data string) string {
	data = strings.Replace(data, "\\u0026", "&", -1)
	data = strings.Replace(data, "\\u003c", "<", -1)
	data = strings.Replace(data, "\\u003e", ">", -1)
	return data
}
