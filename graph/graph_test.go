package main

import (
	//	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	//t.Fatal("not implemented")
}

func TestAddVertex(t *testing.T) {

	var g Graph = Graph{}
	assert.Equal(t, 0, g.GetNoOfVertices())

	var v Vertex = Vertex{}

	assert.Equal(t, true, g.AddVertex(v))
	assert.Equal(t, 1, g.GetNoOfVertices())
}

func TestGetVertexValue(t *testing.T) {

	var g Graph = Graph{}

	var v1 Vertex = Vertex{}
	v1.uuid = "123"
	v1.data = "data1"
	g.AddVertex(v1)

	var v2 Vertex = Vertex{}
	v2.uuid = "245"
	v2.data = "data2"
	g.AddVertex(v2)

	// Retrieve the value for
	var s1 string
	s1 = g.GetVertexValue(Vertex{uuid: "123"})
	assert.Equal(t, "data1", s1)

	// Retrieve the value for
	var s2 string
	s2 = g.GetVertexValue(Vertex{uuid: "245"})
	assert.Equal(t, "data2", s2)

}
