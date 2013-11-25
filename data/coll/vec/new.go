// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package vec

import . "clojang/data/interfaces"

func Vec (seq ISeq) IVector {
  if seq == nil {
    return emptyVector{}
  } else {
    seq = seq.Seq()
    var vec IVector
    vec = emptyVector{}
    for seq != nil {
      vec = vec.Conj(seq.First())
      seq = seq.Rest().Seq()
    }
    
  }
}