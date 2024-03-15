package main

import (
	"fmt"
	"testing"
)

func Test_C(t *testing.T) {
	ma1 := getTgLinkMap()
	for key, val := range ma1 {
		fmt.Println(key, val)
	}
	fmt.Println(ma1["https://www.myswap.xyz/#/"])
}
