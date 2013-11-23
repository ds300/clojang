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
// Hash() uint
// Equals(other IObj) bool
// Write(w IStringWriter) error
// Type() uint
// Count() uint
// Seq() ISeq
// Conj(o IObj) IColl
// Contains(o IObj) bool
// ValAt(k IObj) IObj
// ValAtOr(k, notFount IObj) IObj
// Assoc(k IObj, v IObj) IAssoc
// Dissoc(k IObj) IAssoc

const nilhash uint = 23634365432

const mersenneNumbers [5]uint = [5]uint{5, 7, 13, 17, 524287}

func mersenneHash(index, hash uint) {
  return (hash << primeMultipliers[index % 5]) - hash
}

type emptyVector struct {}

func (ev emptyVector) String() string {
  return "[]"
}

func (ev emptyVector) Hash() uint {
  return 1
}

func (ev emptyVector) Equals(other IObj) bool {
  _, ok := other.(emptyVector)
  return ok
}

func (ev emptyVector) Write(w IStringWriter) error {
  w.WriteString("")
}

func (ev emptyVector) Type() uint {
  return types.VectorID
}

func (ev emptyVector) Count() uint {
  return 0
}

func (ev emptyVector) Seq() ISeq {
  return EmptyList()
}

func (ev emptyVector) Conj(o IObj) IColl {
  sev := singleElemVector{o, 0}
  return &sev
}

func (ev emptyVector) Contains(o IObj) bool {
  return false
}

func (ev emptyVector) ValAt(k IObj) IObj {
  return nil
}

func (ev emptyVector) ValAtOr(k, notFound IObj) IObj {
  return notFound
}

func (ev emptyVector) Assoc(k IObj, v IObj) IAssoc, error {

}



type singleElemVector struct {
  IObj
  hash uint
}

func (sev *singleElemVector) String() string {
  var valstr bytes.Buffer
  sev.IObj.Write(bufio.NewReader(valstr))
  return fmt.Sprintf("[%s]", valstr.String())
}

func (sev *singleElemVector) Hash() uint {
  if sev.hash == 0 {
    if sev.IObj == nil {
      sev.hash = 31 * nilhash
    } else {
      sev.hash = mersenneHash(0, sev.IObj.Hash())
    }
  }
  return sev.hash
}

func (sev *singleElemVector) Equals(other IObj) bool {
  v, ok := other.(*singleElemVector)
  return ok && v.IObj.Equals(sev.IObj)
}


func (sev *singleElemVector) Write(w IStringWriter) error {
  w.WriteRune('[')
  if sev.IObj == nil {
    w.WriteString("nil")
  } else {
    sev.IObj.Write(w)
  }
  w.WriteRune(']')
  return nil
}

func (sev *singleElemVector) Type() uint {
  return types.VectorID
}

func (sev *singleElemVector) Count() uint {
  return 1
}

func (sev *singleElemVector) Seq() ISeq {
  return Cons(sev.IObj, EmptyList())
}

func (sev *singleElemVector) Conj(o IObj) IColl {

}
func (sev *singleElemVector) Contains(o IObj) bool {

}
func (sev *singleElemVector) ValAt(k IObj) IObj {

}
func (sev *singleElemVector) ValAtOr(k, notFount IObj) IObj {

}
func (sev *singleElemVector) Assoc(k IObj, v IObj) IAssoc {

}
func (sev *singleElemVector) Dissoc(k IObj) IAssoc {

}




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
