package coll

import "testing"
import "clojang/data/types/i"
import "fmt"

// we need a thing that implements IObj

type mock struct {
  val interface{}
  hash uint
}

func (m *mock) Equals(o i.IObj) bool {
  v, ok := o.(*mock)
  return ok && v.val == m.val
}

func (m *mock) Hash() uint {
  return m.hash
}

func (m *mock) String() string {
  return fmt.Sprint("(", m.val, ":", m.hash, ")")
}

func newMock(val interface{}, hash uint) *mock {
  m := mock{val, hash}
  return &m
}

func TestMap (t *testing.T) {
  m := NewMap()
  m1 := m.With(newMock("cheese", 323242), newMock("jones", 3234224))

  if m.Contains(newMock("cheese", 323242)) {
    t.Log("immutability breakdown")
    t.Fail()
  }

  if !m1.Contains(newMock("cheese", 323242)) {
    t.Log("the thing didn't get put in")
    t.Fail()
  }

  if m1.Size() != 1 {
    t.Log("adding a thing didn't increase size")
    t.Fail()
  }

  if m1.Hash() != starthash * uint(323242) {
    t.Log("The hash didn't get multiplied")
    t.Fail()
  }
}