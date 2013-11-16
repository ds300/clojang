package types

import "fmt"
import "strconv"

type IObj interface {
  String() string
  Hash() uint
  Equals(other IObj) bool
}

type List struct {
  Val IObj
  Next *List
}

func (list *List) String() string {
  if list.Val == nil {
    return "()"
  } else {
    str := "(" + list.Val.String()
    current := list.Next
    for current.Next != nil {
      str += " " + current.Val.String()
      current = current.Next
    } 
    return str + ")"
  }
}

type Symbol struct {
  Ns string
  Name string
}

func (s *Symbol) String() string {
  if s.Ns == "" {
    return s.Name
  }
  return s.Ns + "/" + s.Name
}

type Keyword struct {
  Ns string
  Name string
}

func (k *Keyword) String() string {
  if k.Ns == "" {
    return ":" + k.Name
  }
  return ":" + k.Ns + "/" + k.Name
}

type Long int64

func (l *Long) String() string {
  return fmt.Sprintf("%d", l)
}

type Double float64

func (d *Double) String() string {
  return fmt.Sprintf("%g", *d)
}

type Fraction struct {
  Nom int64
  Denom int64
}

func (f *Fraction) String() string {
  return fmt.Sprintf("%d/%d", f.Nom, f.Denom)
}

type String string

func (s *String) String() string {
  return strconv.Quote(string(*s))
}

type Bool bool

func (b *Bool) String() string {
  if bool(*b) {
    return "true"
  } else {
    return "false"
  }
}

type Nil bool

func (n *Nil) String() string {
  return "nil"
}

func (n *Nil) Equal() bool {

}