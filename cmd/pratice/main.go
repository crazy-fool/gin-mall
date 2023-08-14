package main

import (
	"fmt"
	"gin-mall/pkg/helper/sign"
)

func main() {
	sign := sign.NewRsaSign("", "")
	params := map[string]any{
		"a": 123,
		"b": "ccc",
		"s": "12.3",
	}
	str := sign.SortParams(params)
	fmt.Println(str)
}
