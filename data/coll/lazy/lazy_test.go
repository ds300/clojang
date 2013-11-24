package lazy

import "testing"
import . "clojang/data/interfaces"
import "clojang/data/coll/list"
import "clojang/data/primitives"

func take (n uint, seq ISeq) ISeq {
  if n==0 || seq == nil {
    return nil
  } else {
    return LazySeq(func () ISeq {
      return list.Cons(seq.First(), take(n-1, seq.Rest()))
    })
  }
}


func naturalIntegers (from uint) ISeq {
  return LazySeq(func () ISeq {
    return list.Cons(primitives.Long(from), naturalIntegers(from+1))
  })
}

func TestLazy (t *testing.T) {
  t.Log("ok")



  seq := take(10, naturalIntegers(0))

  head := seq

  for seq.Seq() != nil {
    t.Log(seq.First())
    seq = seq.Rest()
  }

  t.Log("That first value again was", head.First())

  t.Log(head)
}