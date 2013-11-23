// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package sequtil

import . "clojang/data/interfaces"

func WriteSeq(w IStringWriter, leftDelim rune, seq ISeq, rightDelim rune) error {
  _, err := w.WriteRune(leftDelim)

  seq = seq.Seq()
  for err == nil && seq != nil {
    val := seq.First()
    if val == nil {
      _, err = w.WriteString("nil")
    } else {
      err = val.Write(w)
    }

    seq = seq.Rest().Seq()
    if err == nil && seq != nil {
      _, err = w.WriteRune(' ')
    }
  }

  _, err = w.WriteRune(rightDelim)
  return err
}