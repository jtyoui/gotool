// Package math
// @Time  : 2022/7/11 13:15
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package math

import (
	"math/rand"
	"reflect"
)

// RandomRange returns a random int in [min, max).
func RandomRange[T int | int32 | int64](min, max T) T {
	typ := reflect.TypeOf(max)
	switch typ.Kind() {
	case reflect.Int:
		t := int(min) + rand.Intn(int(max-min))
		return T(t)
	case reflect.Int32:
		t := int32(min) + rand.Int31n(int32(max-min))
		return T(t)
	case reflect.Int64:
		t := int64(min) + rand.Int63n(int64(max-min))
		return T(t)
	}
	return T(0)
}
