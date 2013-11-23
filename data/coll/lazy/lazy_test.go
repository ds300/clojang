package coll

import "testing"
import . "clojang/data/interfaces"

func take (n uint, seq ISeq) ISeq {
  if n==0 || seq == nil {
    return nil
  } else {
    return LazySeq(func () ISeq {
      return Cons(seq.First(), take(n-1, seq.Rest()))
    })
  }
}


func naturalIntegers (from uint) ISeq {
  return LazySeq(func () ISeq {
    return Cons(mock(from, from), naturalIntegers(from+1))
  })
}

func TestLazy (t *testing.T) {
  t.Log("ok")



  seq := take(10, naturalIntegers(0))

  head := seq

  for seq != nil {
    t.Log(seq.First())
    seq = seq.Rest()
  }

  t.Log("That first value again was", head.First())

  t.Log(head)
}