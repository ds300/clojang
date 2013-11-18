package coll

import "clojang/data/i"

type sliceVector []i.IObj

func sliceVectorSeq(v *sliceVector, i, n int) i.ISeq {
  return LazySeq(func {} i.ISeq {
    if i < n {
      return Cons(v[i], sliceVectorSeq(v, i+1, n))
    } else {
      return nil
    }
  })
}



func (v *sliceVector) Seq() i.ISeq {
  i := 0
  n := len(*v)
  return LazySeq(func () i.ISeq {
    if i < n {
      return Cons()
    }
    })
}