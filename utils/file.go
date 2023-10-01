package utils

import "strings"

// TODO: 修复文件中含有\
func CheckFileVaild(f string) string {
	fileNameInvaildcharacter := [...]string{"<", ">", "/", "//", "|", ":", "*", "?", "\n"}
	for _, c := range fileNameInvaildcharacter {
		// 如果文件名中含有该字符则将文件名中所有该字符替换为 '-'
		f = strings.Replace(f, string(c), "_", -1)
	}
	return f
}
