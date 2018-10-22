package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSizeOfEmptyStack(t *testing.T) {

  s := data_structure_int { data: make([]int, 5) }
  assert.Equal(t, 0, s.Size(), "The size of an empty stack should be 0")

}

func TestPush (t *testing.T) {

  s := data_structure_int { data: make([]int, 5) }

  assert.Equal(t, s.Push(2), true, "Push should return true" )
  assert.Equal(t, 1, s.Size(), "Size of the stack should be 1")
}

func TestPop (t *testing.T) {

  s := data_structure_int { data: make([]int, 5) }

  s.Push(1);
  s.Push(2);
  s.Push(3);

  assert.Equal(t, 3, s.Pop(), "Expected 3 to be popped" )
  assert.Equal(t, 2, s.Size(), "Stack length should be 2")

  assert.Equal(t, 2, s.Pop(),  "Expected 2 to be popped" )
  assert.Equal(t, 1, s.Size(), "Stack length should be 1")

  assert.Equal(t, 1, s.Pop(),  "Expected 1 to be popped" )
  assert.Equal(t, 0, s.Size(), "Stack length should be 0")


}

func TestGrowthOfStack (t *testing.T) {

  s := data_structure_int { data: make([]int, 5) }
  s.Push(1);
  s.Push(2);
  s.Push(3);
  s.Push(4);
  s.Push(5);
  s.Push(6);
  assert.Equal(t, 6, s.Size(), "Stack length should be 6")

  s.Push(7);
  assert.Equal(t, 7, s.Size(), "Stack length should be 7")

  s.Pop()
  assert.Equal(t, 6, s.Size(), "Stack length should be 6")

  s.Pop()
  assert.Equal(t, 5, s.Size(), "Stack length should be 5")

}
