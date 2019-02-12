package main

import (
//"fmt"
)

type Graph struct {
	vertices       []Vertex
	no_of_vertices int
}

type Vertex struct {
	uuid string
	data string
}

func (g *Graph) GetNoOfVertices() int {
	return g.no_of_vertices
}

func (g *Graph) AddVertex(v Vertex) bool {
	g.vertices = append(g.vertices, v)
	g.no_of_vertices = len(g.vertices)
	return true
}

func (g *Graph) GetVertexValue(v Vertex) string {

	for _, stored_vertex := range g.vertices {
		if stored_vertex.uuid == v.uuid {
			return stored_vertex.data
		}
	}

	return ""

}
