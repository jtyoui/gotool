// Package excel
// @Time  : 2022/7/3 22:17
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel_test

import (
	"fmt"
	"github.com/jtyoui/gotool/file/excel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveExcel(t *testing.T) {
	values := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	err := excel.SaveExcel("test.xlsx", values)
	assert.NoError(t, err)

	data, _ := excel.LoadExcel[Test]("test.xlsx")
	assert.Equal(t, data, values)

	values1 := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	err1 := excel.SaveExcel("test.xlsx", values1)
	assert.NoError(t, err1)

	data1, _ := excel.LoadExcel[Test]("test.xlsx")
	assert.Equal(t, data1, values1)
}

func ExampleSaveExcel() {
	/***
	type Test struct {
		Name  string `excel:"name"`
		Age   int    `excel:"age"`
		Sex   string `excel:"sex"`
		High  int    `excel:"-"`
		Width int
	}

	func (t Test) GetXLSXSheetName() string {
		return "Sheet1"
	}
	*/

	values := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/
	err := excel.SaveExcel("test.xlsx", values)
	fmt.Println(err)
	// Output:
	// <nil>
}
