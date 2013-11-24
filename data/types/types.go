// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package types
// import "fmt"
// import "strconv"

const (
  BoolID uint32 = iota
  LongID uint32 = iota
  DoubleID uint32 = iota
  StringID uint32 = iota
  KeywordID uint32 = iota
  SymbolID uint32 = iota
  ListID uint32 = iota
  VectorID uint32 = iota
  MapID uint32 = iota 
  SetID uint32 = iota
  RegexID uint32 = iota
  LazySeqID uint32 = iota
)


// type List struct {
//   Val IObj
//   Next *List
// }

// func (list *List) String() string {
//   if list.Val == nil {
//     return "()"
//   } else {
//     str := "(" + list.Val.String()
//     current := list.Next
//     for current.Next != nil {
//       str += " " + current.Val.String()
//       current = current.Next
//     } 
//     return str + ")"
//   }
// }

// type Symbol struct {
//   Ns string
//   Name string
// }

// func (s *Symbol) String() string {
//   if s.Ns == "" {
//     return s.Name
//   }
//   return s.Ns + "/" + s.Name
// }

// type Keyword struct {
//   Ns string
//   Name string
// }

// func (k *Keyword) String() string {
//   if k.Ns == "" {
//     return ":" + k.Name
//   }
//   return ":" + k.Ns + "/" + k.Name
// }

// type Long int64

// func (l *Long) String() string {
//   return fmt.Sprintf("%d", l)
// }

// type Double float64

// func (d *Double) String() string {
//   return fmt.Sprintf("%g", *d)
// }

// type Fraction struct {
//   Nom int64
//   Denom int64
// }

// func (f *Fraction) String() string {
//   return fmt.Sprintf("%d/%d", f.Nom, f.Denom)
// }

// type String string

// func (s *String) String() string {
//   return strconv.Quote(string(*s))
// }

// type Bool bool

// func (b *Bool) String() string {
//   if bool(*b) {
//     return "true"
//   } else {
//     return "false"
//   }
// }

// type Nil bool

// func (n *Nil) String() string {
//   return "nil"
// }

// func (n *Nil) Equal() bool {

// }