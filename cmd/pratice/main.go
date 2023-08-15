package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//sign := sign.NewRsaSign("", "")
	//params := map[string]any{
	//	"a": 123,
	//	"b": "ccc",
	//	"s": "12.3",
	//}
	//str := sign.SortParams(params)
	//fmt.Println(str)
	var c any
	c = 1
	str := convertToString(c)
	fmt.Println(str)
}

func a(i any) {
	v, ok := i.(string)
	fmt.Println(v, ok)
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
