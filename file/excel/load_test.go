// Package excel
// @Time  : 2022/7/3 9:24
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel_test

import (
	"fmt"
	"github.com/jtyoui/gotool/file/excel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadXlsx(t *testing.T) {
	data, err := excel.LoadExcel[Test]("test.xlsx")
	assert.NoError(t, err)
	test1 := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	assert.Equal(t, data, test1)
}

func ExampleLoadExcel() {
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

	// Excel File view
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/
	data, _ := excel.LoadExcel[Test]("test.xlsx")
	fmt.Println(data)
	// Output:
	// [{张三 17 男 0 0} {李四 18 女 0 0}]
}
