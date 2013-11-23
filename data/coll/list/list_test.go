package list

import . "clojang/data/interfaces"
import "testing"
import "clojang/data/primitives"

func TestList(t *testing.T) {
  var ls ISeq
  for i := 0; i < 10; i++ {
    ls = Cons(primitives.Long(i), ls)
  }
  t.Log(ls)
  t.Log(ls.(ICounted).Count())
  t.Log(ls.(IObj).Hash())
}