package main

import "fmt"

type FinishedHouse struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Option func(*FinishedHouse)

func NewFinishedHouse(options ...Option) *FinishedHouse {
	h := &FinishedHouse{
		// default options
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}
	for _, option := range options {
		option(h)
	}

	return h
}

func WithStyle(style int) Option {
	return func(h *FinishedHouse) {
		h.style = style
	}
}
func WithCentralAirConditioning(centralAirConditioning bool) Option {
	return func(h *FinishedHouse) {
		h.centralAirConditioning = centralAirConditioning
	}
}

func WithFloorMaterial(floorMaterial string) Option {
	return func(h *FinishedHouse) {
		h.floorMaterial = floorMaterial
	}
}

func WithWallMaterial(wallMaterial string) Option {
	return func(h *FinishedHouse) {
		h.wallMaterial = wallMaterial
	}
}

func main() {
	// default
	fmt.Printf("%+v\n", NewFinishedHouse())
	fmt.Printf("%+v\n", NewFinishedHouse(WithStyle(1),
		WithCentralAirConditioning(false),
		WithFloorMaterial("ground-title")))
}
