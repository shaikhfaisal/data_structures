package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	//t.Fatal("not implemented")
}

func TestAddVertex(t *testing.T) {

	var g Graph = Graph{}
	var v Vertex = Vertex{}

	assert(t, true, g.AddVertex(v))
}
