package util

import "fmt"
import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}


func ToString[T Number](v T) string  {
	return fmt.Sprint(v)
}
