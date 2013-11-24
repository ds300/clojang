// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package vec

import "fmt"
import . "clojang/data/interfaces"
import "clojang/data/types"
import "bufio"
import "bytes"

// String() string
// Hash() uint32
// Equals(other IObj) bool
// Write(w IStringWriter) error
// Type() uint32
// Count() uint32
// Seq() ISeq
// RSeq() ISeq
// Conj(o IObj) IColl
// Assoc(k IObj, v IObj) (IAssoc, error)
// Get(k IObj) IObj
// GetOr(k, notFound IObj) IObj
// Contains(o IObj) bool


type singleElemVector struct {
  IObj
  hash uint
}


