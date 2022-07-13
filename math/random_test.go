// Package math
// @Time  : 2022/7/13 15:26
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package math_test

import (
	"fmt"
	"github.com/jtyoui/gotool/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomRange(t *testing.T) {
	min, max := 5, 10
	value := math.RandomRange(min, max)
	assert.True(t, value >= min && value <= max)

	min1, max1 := int32(1021), int32(10545121)
	value1 := math.RandomRange(min1, max1)
	assert.True(t, value1 >= min1 && value1 <= max1)

	min2, max2 := int64(10216545121), int64(105451211564612313)
	value2 := math.RandomRange(min2, max2)
	assert.True(t, value2 >= min2 && value2 <= max2)
}

func ExampleRandomRange() {
	min, max := 9, 10
	value := math.RandomRange(min, max)
	fmt.Println(value)
	// Output:
	// 9
}
