// Copyright (c) David Sheldrick. All rights reserved.
// The use and distribution terms for this software are covered by the
// Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
// which can be found in the file epl-v10.html at the root of this distribution.
// By using this software in any fashion, you are agreeing to be bound by
// the terms of this license.
// You must not remove this notice, or any other, from this software.

package main

import "fmt"
import "clojang/data/types"
// import "os"
// import "clojang/data/parse"

func main () {
  // if len(os.Args) == 1 {
  //   k := types.Keyword{"clojang.core", "key"}
  //   ls := types.List{&k, nil}
  //   fmt.Println(ls.String())

  // } else {
  //   fmt.Println(parse.ReadFile(os.Args[1]))
  // }
  types.IPopcount(3, 5)
  fmt.Println(types.Popcount(0xFFFF0000))
}