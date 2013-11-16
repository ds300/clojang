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