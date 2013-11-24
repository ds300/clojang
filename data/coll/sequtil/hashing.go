// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package sequtil

import . "clojang/data/interfaces"

var mersenneNumbers [5]uint32 = [5]uint32{5, 7, 13, 17, 19}

func HashIndexed(index, hash uint32) uint32 {
  return ((hash << mersenneNumbers[index % 5]) - hash)
}

const nilhash = 329837298

func HashSeq(seq ISeq) uint32 {
  i := uint32(0)

  hash := uint32(5843)

  for seq.Seq() != nil {
    val := seq.First()
    if val == nil {
      hash += HashIndexed(i, nilhash)
    } else {
      hash += HashIndexed(i, val.Hash())
    }
    i++
    seq = seq.Rest()
  }

  if hash == 0 {
    hash = 1
  }

  return hash
}