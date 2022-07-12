// Package excel
// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel

import (
	"errors"
	"fmt"
	"github.com/jtyoui/gotool/str"
	"github.com/xuri/excelize/v2"
	"reflect"
)

// Xlsx loading Excel must be implementation interface.
type Xlsx interface {
	// GetXLSXSheetName return need loading sheet name in the Excel.
	GetXLSXSheetName() string
}

// LoadExcel loading Excel file
//
// must be implemented Xlsx interface.
func LoadExcel[T Xlsx](filePath string) (data []T, err error) {
	// open Excel file
	var file *excelize.File
	if file, err = excelize.OpenFile(filePath); err != nil {
		return
	}

	defer func() {
		// Close the spreadsheet.
		if err = file.Close(); err != nil {
			return
		}
	}()

	// get sheet name from Xlsx interface
	sheetName := T.GetXLSXSheetName(*new(T))

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return
	}

	data = make([]T, len(rows)-1)

	// get title row
	title := map[string]int{}
	for i, cell := range rows[0] {
		title[cell] = i
	}

	// get data row
	for index, row := range rows[1:] {
		var t T
		value := reflect.ValueOf(&t)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			tag := field.Tag.Get("excel")
			if tag == "" || tag == "-" {
				continue
			}

			if j, ok := title[tag]; !ok {
				err = fmt.Errorf("%s is not in the title", tag)
				return
			} else {
				v := value.Field(i)
				if err = cp(&v, row[j]); err != nil {
					return
				}
			}
		}
		data[index] = t
	}
	return
}

// cp copy value from string to value.
func cp(v1 *reflect.Value, b string) error {
	switch v1.Kind() {
	case reflect.Bool:
		if v2, err := str.To[bool](b); err != nil {
			return err
		} else {
			v1.SetBool(v2)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v2, err := str.To[int64](b); err != nil {
			return err
		} else {
			v1.SetInt(v2)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v2, err := str.To[uint64](b); err != nil {
			return err
		} else {
			v1.SetUint(v2)
		}
	case reflect.Float32, reflect.Float64:
		if v2, err := str.To[float64](b); err != nil {
			return err
		} else {
			v1.SetFloat(v2)
		}
	case reflect.String:
		v1.SetString(b)
	default:
		return errors.New("unsupported type: " + v1.Type().String())
	}
	return nil
}
