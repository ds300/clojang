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

type Double float64

func (d Double) String() string {
  return fmt.Sprint(float64(d))
}

func (d Double) Hash() uint32 {
  return NewString(d.String()).Hash()
}

func (d Double) Equals(other IObj) bool {
  switch other.(type) {
  case Long:
    return float64(d) == float64(other.(Long))
  case Double:
    return d == other.(Double)
  default:
    return false
  }
}

func (d Double) Write(w IStringWriter) error {
  _, err := w.WriteString(d.String())
  return err
}

func (d Double) Type() uint32 {
  return types.DoubleID
}

func (d Double) Mult(other INumeric) INumeric {
  switch other.(type) {
  case Long:
    return Double(float64(d) * float64(other.(Long)))
  case Double:
    return Double(d * other.(Double))
  default:
    panic("oh no")
  }
}

func (d Double) Div(other INumeric) (INumeric, error) {
  switch other.(type) {
  case Long:
    if other.(Long) == 0 {
      return nil, errors.New("Divide by zero")
    } else {
      return Double(float64(d) / float64(other.(Long))), nil
    }
  case Double:
    if other.(Double) == 0 {
      return nil, errors.New("Divide by zero")
    } else {
      return Double(d / other.(Double)), nil
    }
  default:
    panic("oh no")
  }
}

func (d Double) Plus(other INumeric) INumeric {
  switch other.(type) {
  case Long:
    return Double(float64(other.(Long)) + float64(d))
  case Double:
    return Double(other.(Double) + d)
  default:
    panic("oh no")
  }
}

func (d Double) Sub(other INumeric) INumeric {
  switch other.(type) {
  case Long:
    return Double(float64(d) - float64(other.(Long)))
  case Double:
    return Double(d - other.(Double))
  default:
    panic("oh no")
  }
}

func (d Double) Mod(other INumeric) (INumeric, error) {
    return nil, errors.New("Illegal operand type for mod: double")
}
