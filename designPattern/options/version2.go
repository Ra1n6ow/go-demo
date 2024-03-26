package main

import (
	"fmt"
)

/*
使用结构体封装配置选项,该构造函数的api签名可保持不变
*/

type FinishedHouse2 struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Options2 struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

func NewFinishedHouse2(options *Options2) *FinishedHouse2 {
	var style int = 0
	var centralAirConditioning bool = true
	var floorMaterial string = "wood"
	var wallMaterial string = "paper"

	if options != nil {
		style = options.style
		centralAirConditioning = options.centralAirConditioning
		floorMaterial = options.floorMaterial
		wallMaterial = options.wallMaterial
	}

	h := &FinishedHouse2{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           wallMaterial,
	}
	return h
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse2(nil))
	o := &Options2{
		style:                  1,
		centralAirConditioning: false,
		floorMaterial:          "ground-title",
		wallMaterial:           "paper",
	}
	fmt.Printf("%+v\n", NewFinishedHouse2(o))
}
