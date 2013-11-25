package vec

import "testing"
import "clojang/data/primitives"
import . "clojang/data/interfaces"


func TestVec(t *testing.T) {
  var x IObj = emptyVector{}

  getNil(t, x.(IVector), -2, -1, 0, 1, 2, 42374827, -34253465)

  x = x.(IVector).Conj(primitives.Long(43)).(IObj)

  t.Log("The single vector looks like this: ", x.String())

  x1, err := x.(IVector).Assoc(primitives.Long(0), primitives.Long(52))

  t.Log("You can make a new one with a different value: ", x1.(IObj).String())

  if err != nil {
    t.Fail()
    t.Log("You shouldn't get an error when you do that")
  }

  x1, err = x.(IVector).Assoc(primitives.Long(2), primitives.Long(342))

  if err == nil {
    t.Fail()
    t.Log("there should be an index out of bounds error")
  }

  t.Log("trying to assoc to an out-of-bounds index gives an error like:", err)

  

  x = x.(IVector).Conj(primitives.NewString("Cheese")).(IObj)

  t.Log("Tuple vectors look like this: ", x.String())


  x = x.(IVector).Conj(primitives.NewString("yo")).(IObj)

  t.Log("Slice vectors look like this: ", x.String())

  x = x.(IVector).Conj(primitives.NewString(`yo again!
    `)).(IObj)

  t.Log("Slice vectors grow man: ", x.String())
}

func getNil(t *testing.T, vec IVector, indices ...int) {
  notFound := primitives.NewString("it wasn't found!")
  for i := range indices {
    if vec.Get(primitives.Long(i)) != nil {
      t.Fail()
      t.Log("Getting should return nill for index", i, "in vector", vec.(IObj).String())
    }
    if vec.GetOr(primitives.Long(i), notFound) != notFound {
      t.Fail()
      t.Log("Getting should return notFound for index", i, "in vector", vec.(IObj).String())
    }
  }
}