package main

import (
	"math/rand"
)

type id struct {
	value int
}

func newId() id {
	return id{
		value: rand.Int(),
	}
}
