package main

import "fmt"

/*
最致命的是该接口无法扩展。如果我们此时应用户要求增加一个室内门型设置的选项（可选实木门或板材套装门），该接口无法满足。
考虑兼容性原则，该接口一旦发布就成为API的一部分，我们不能随意变更。
*/

type FinishedHouse1 struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

func NewFinishedHouse1(style int, centralAirConditioning bool, floorMaterial, wallMaterial string) *FinishedHouse1 {
	h := &FinishedHouse1{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           floorMaterial,
	}
	return h
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse1(0, true, "wood", "paper"))
}
