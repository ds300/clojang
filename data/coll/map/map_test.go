package coll

import "testing"


// we need a thing that implements IObj


func TestMap (t *testing.T) {
  m := NewMap()
  m1 := m.With(mock("cheese", 323242), mock("jones", 3234224))

  if m.Contains(mock("cheese", 323242)) {
    t.Log("immutability breakdown")
    t.Fail()
  }

  if !m1.Contains(mock("cheese", 323242)) {
    t.Log("the thing didn't get put in")
    t.Fail()
  }

  if !m1.EntryAt(mock("cheese", 323242)).Val.Equals(mock("jones", 3)) {
    t.Log("the thing got put in but isn't being returned")
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