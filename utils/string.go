package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ConcatXBogusUrlString(reqStruct interface{}, urlPrefix string) string {
	var url = strings.Builder{}
	url.WriteString(urlPrefix)
	typ := reflect.TypeOf(reqStruct)
	val := reflect.ValueOf(reqStruct)
	num := val.Elem().NumField()
	for i := 0; i < num; i++ {
		val := fmt.Sprint(val.Elem().Field(i))
		tagVal := typ.Elem().Field(i).Tag.Get("json")
		url.WriteString(tagVal + "=" + val + "&")
	}
	return url.String()[:len(url.String())-1]
}

func PrintInitLogo() {
	fmt.Println("欢迎使用抖音视频下载器, 关注请访问https://github.com/root-wang/TikTokDownloadVideo")
	w := ".___. .  .___.   .  		.__        ,     .  .         \n" +
		"  |  *;_/  |   _ ;_/		[__) _  _ -+-    |  | _.._  _ \n" +
		"  |  || \\  |  (_)| \\		|  \\(_)(_) | ____|/\\|(_][ )(_]\n" +
		"                    			                   ._|"
	fmt.Println("\n", w)
}
