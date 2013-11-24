package sequtil

import . "clojang/data/interfaces"

func Equals (a, b ISeq) bool {
  if a != nil && b != nil {
    seqa := a.Seq()
    seqb := b.Seq()
    for seqa != nil && seqb != nil {
      if !seqa.First().Equals(seqb.First()) {
        return false
      }
    }
    return seqa == nil && seqb == nil
  } else {
    return a == nil && b == nil
  }
}