package main

import "fmt"

func main() {
  fmt.Printf("Nothing to see here. Its all about the tests :) \n")
}

type data_structure_int struct {
  data []int
  index int
}

type stack_int interface {
  Push() bool
  pop() int
  Size() int
}

func (ds *data_structure_int) Push(element int) bool {

  if (ds.index >= len(ds.data)) {
    ds.grow()
  }

  ds.data[ds.index] = element
  ds.index++

  return true
}

func (ds *data_structure_int) Pop() int {
  ds.index--
  return ds.data[ds.index]
}

func (ds *data_structure_int) Size() int {
  return ds.index
}

func (ds *data_structure_int) grow() bool {

  new_data := make ( []int, len(ds.data)*2)
  copy(new_data, ds.data)
  ds.data = new_data
  return true
}

