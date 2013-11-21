package coll

import . "clojang/data/interfaces"
import "fmt"
import "bufio"

// we need a thing that implements IObj

type mockIObj struct {
  val interface{}
  hash uint
}


func (m *mockIObj) Write(w bufio.Writer) {
  w.WriteString(m.String())
}


func (m *mockIObj) Equals(o IObj) bool {
  v, ok := o.(*mockIObj)
  return ok && v.val == m.val
}

func (m *mockIObj) Hash() uint {
  return m.hash
}

func (m *mockIObj) String() string {
  return fmt.Sprint("(", m.val, ":", m.hash, ")")
}

func mock(val interface{}, hash uint) *mockIObj {
  m := mockIObj{val, hash}
  return &m
}