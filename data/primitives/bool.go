// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package primitives

import . "clojang/data/interfaces"
import "clojang/data/types"
import "bufio"

type Bool bool


func (b Bool) String() string {
  if b {
    return "true"
  } else {
    return "false"
  }
}

func (b Bool) Hash() uint32 {
  if b {
    return 2
  } else {
    return 1
  }
}

func (b Bool) Equals(other IObj) bool {
  v, ok := other.(Bool)
  return ok && v == b
}

func (b Bool) Write(w *bufio.Writer) error {
  _, err := w.WriteString(b.String())
  return err
}

func (b Bool) Type() uint32 {
  return types.BoolID
}
