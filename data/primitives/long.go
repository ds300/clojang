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
import "fmt"
import "errors"
import "bufio"

type Long int64

func (l Long) String() string {
  return fmt.Sprint(int64(l))
}

func (l Long) Hash() uint32 {
  return uint32(l)
}

func (l Long) Equals(other IObj) bool {
  switch other.(type) {
  case Long:
    return l == other.(Long)
  case Double:
    return float64(l) == float64(other.(Double))
  default:
    return false
  }
}

func (l Long) Write(w *bufio.Writer) error {
  _, err := w.WriteString(l.String())
  return err
}

func (l Long) Type() uint32 {
  return types.LongID
}

func (l Long) Mult(other INumeric) INumeric {
  switch other.(type) {
  case Long:
    return Long(l * other.(Long))
  case Double:
    return Double(Double(l) * other.(Double))
  default:
    panic("oh no")
  }
}

func (l Long) Div(other INumeric) (INumeric, error) {
  switch other.(type) {
  case Long:
    if other.(Long) == 0 {
      return nil, errors.New("Divide by zero")
    } else {
      return Long(l / other.(Long)), nil
    }
  case Double:
    if other.(Double) == 0 {
      return nil, errors.New("Divide by zero")
    } else {
      return Double(Double(l) / other.(Double)), nil
    }
  default:
    panic("oh no")
  }
}

func (l Long) Plus(other INumeric) INumeric {
  switch other.(type) {
  case Long:
    return Long(other.(Long) + l)
  case Double:
    return Double(float64(other.(Double)) + float64(l))
  default:
    panic("oh no")
  }
}

func (l Long) Sub(other INumeric) INumeric {
  switch other.(type) {
  case Long:
    return Long(l - other.(Long))
  case Double:
    return Double(Double(l) - other.(Double))
  default:
    panic("oh no")
  }
}

func (l Long) Mod(other INumeric) (INumeric, error) {
  switch other.(type) {
  case Long:
    return Long(l % other.(Long)), nil
  case Double:
    return nil, errors.New("Illegal mod arguments: Long, Double")
  default:
    panic("oh no")
  }
}
