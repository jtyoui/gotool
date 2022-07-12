// Package excel
// @Time  : 2022/7/12 9:20
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel_test

type Test struct {
	Name  string `excel:"name"`
	Age   int    `excel:"age"`
	Sex   string `excel:"sex"`
	High  int    `excel:"-"`
	Width int
}

func (Test) GetXLSXSheetName() string {
	return "test"
}
