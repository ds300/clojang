// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

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
  return sliceVectorSeq(v, 0, len(*v))
}
