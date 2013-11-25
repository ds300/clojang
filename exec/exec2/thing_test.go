package exec2

import "testing"
import . "clojang/data/coll/list"
import . "clojang/data/primitives"

func TestInterpret (t *testing.T) {
  instrs := []interface{}{
    PUSH{Cons(Long(4), Cons(Long(5), nil))},
    CALL{&stackFrame{"plus", [2]int{0,0}, 1, nil}, &plus},
    PUSH{Cons(Long(4), Cons(Long(5), nil))},
    CALL{&stackFrame{"print", [2]int{0,0}, 1, nil}, &print},
    RETURN{},
  }
  interpreter(instrs)(nil, EmptyList{})
  // instrs := []byte{PUSH, PUSH, ADD, PRINT}
  // things := [][]int{[]int{5}, []int{4}, nil, nil}
  
  // for i := 0; i < 20; i++ {
  //   instrs = append(instrs, instrs...)
  //   things = append(things, things...)
  // }

  // instrs = append(instrs, RETURN)
  // things = append(things, nil)

  // interpreter(instrs, things)
  // t.Log("Damn I wrote", len(things), "9s")
}