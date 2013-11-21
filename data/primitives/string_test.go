package primitives

import "testing"

func TestString(t *testing.T) {
  s1 := NewString("hello")
  s2 := NewString("goodbye")
  t.Log("s1 hash:", s1.Hash())
  t.Log("s2 hash:", s2.Hash())

  s3 := NewString("he" + "llo")
  if !s1.Equals(s3) {
    t.Log("Strings which should be equal are not equal", s1, s3)
    t.Fail()
  }
}