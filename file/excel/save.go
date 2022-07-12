// Package excel
// @Time  : 2022/7/3 10:12
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

// SaveExcel save data to excel.
//
// must be implemented Xlsx interface.
func SaveExcel[T ~[]E, E Xlsx](filepath string, data T) (err error) {
	xlsx := excelize.NewFile()
	sheet := data[0].GetXLSXSheetName()
	index := xlsx.NewSheet(sheet)

	s := reflect.ValueOf(data)

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i)
		// drop ptr
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		elemType := elem.Type()

		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("excel")
			if tag == "" || tag == "-" {
				continue
			}
			column := 'A' + j
			if i == 0 {
				err = xlsx.SetCellValue(sheet, fmt.Sprintf("%c%d", column, i+1), tag)
			}

			err = xlsx.SetCellValue(sheet, fmt.Sprintf("%c%d", column, i+2), elem.Field(j).Interface())
		}
	}
	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	err = xlsx.SaveAs(filepath)
	return
}
