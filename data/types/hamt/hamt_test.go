package hamt

import "testing"
import "fmt"
import "clojang/data/types/i"

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

// this one is pretty trivial
func TestEmpty (t *testing.T) {
  empty := new(EmptyNode)

  entry := NewEntry(newMock("hello", 5), newMock("world", 6))

  empty2, incCount := empty.With(entry, 5, 0)

  if empty2 != entry || !incCount {
    t.Log("assertion failed: empty2 == entry && incCount")
    t.Fail()
  }

  empty3, decCount := empty.Without(newMock("hello", 5), 5, 0)

  if empty3 != empty || decCount {
    t.Log("assertion failed: empty3 == empty && !decCount")
    t.Fail()
  }

  val := empty.EntryAt(newMock("hello", 5), 5, 0)

  if val != nil {
    t.Log("Assertion failed: val == nil")
    t.Fail()
  }
}

// this one is also pretty trivial
func TestEntry (t *testing.T) {
  entry := NewEntry(newMock("hello", 5), newMock("world", 6))

  obj := entry.EntryAt(newMock("hello", 5), 5, 0)

  if !obj.Val.Equals(newMock("world", 6)) {
    t.Log("Entry returned the wrong thing")
    t.Fail()
  }

  empty, decCount := entry.Without(newMock("hello", 5), 5, 0)

  if empty != nil || !decCount {
    t.Log("Entry.Without didn't return nil")
    t.Fail()
  }

  entry2, incCount := entry.With(NewEntry(newMock("collide", 5), newMock("crash", 10)), 5, 0)

  _, ok := entry2.(*collisionNode)

  if !ok {
    t.Log("Entry.With didn't return a collision node when we made a collision")
    t.Fail()
  }

  if !incCount {
    t.Log("Entry.With didn't cause count increase on collision")
    t.Fail()
  }

  entry3, incCount2 := entry.With(NewEntry(newMock("partialCollide", 69), newMock("ouch", 10)), 69, 0)

  _, ok2 := entry3.(*hamtNode)

  if !ok2 {
    t.Log("Entry.With for partial collide didn't create distinguishing node")
    t.Fail()
  }

  if !incCount2 {
    t.Log("Entry.With for partial collide didn't increase count")
    t.Fail()
  }
}

func TestDistinguish (t *testing.T) {
  var n1 uint = 0x3FFFFFFF // 0b0011111.... so 0 when we get to the last level
  var n2 uint = 0x7FFFFFFF // 0b01111111... so 1 when we get to the last level
  e1 := NewEntry(newMock("hello", n1), newMock("yo", 4))
  e2 := NewEntry(newMock("hey", n2), newMock("there", 5))

  // this should create a tree 6 levels deep
  // top level for 1st 5 bits
  // 1st level for 2nd 5 bits
  // ...
  // 5th level for 6th 5 bits
  // 6th level for last 2 bits
  hn := distinguishingNode(e1, e2, n1, n2, 0)

  for i:=0;i<6;i++ {
    if len(hn.kids) != 1 {
      t.Log("unexpected nodes were created")
      t.Fail()
    }
    hn = hn.kids[0].(*hamtNode)
  }

  if hn.kids[0] != e1 {
    t.Log("e1 was not in the suspected place")
    t.Fail()
  }

  if hn.kids[0] != e1 {
    t.Log("e1 was not in the suspected place")
    t.Fail()
  }
}



