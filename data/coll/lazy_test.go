package coll

import "testing"
import "clojang/data/i"

func take (n uint, seq i.ISeq) i.ISeq {
  if n==0 || seq == nil {
    return nil
  } else {
    return LazySeq(func () i.ISeq {
      return Cons(seq.First(), take(n-1, seq.Rest()))
    })
  }
}


func naturalIntegers (from uint) i.ISeq {
  return LazySeq(func () i.ISeq {
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