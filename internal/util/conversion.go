package util

import "fmt"

func ToString[T any](v T) string  {
	return fmt.Sprint(v)
}
