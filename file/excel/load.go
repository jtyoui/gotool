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
	"strings"
)

// Excel loading Excel must be implementation interface.
type Excel interface {
	// GetSheetName return need loading sheet name in the Excel.
	GetSheetName() string
}

// LoadExcel loading Excel file
//
// must be implemented Excel interface.
func LoadExcel[T Excel](filePath string) (data []T, err error) {
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

	Object := reflect.TypeOf(data).Elem()
	if Object.Kind() == reflect.Ptr {
		Object = Object.Elem()
	}

	// get sheet name from Excel interface
	sheetName := reflect.New(Object).Interface().(T).GetSheetName()

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
		t := reflect.New(Object).Interface().(T)

		value := reflect.ValueOf(&t)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		value = reflect.Indirect(value)

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			tag := field.Tag.Get("excel")
			if tag == "" || tag == "-" {
				continue
			}

			// get split sep for tag
			var split string
			if strings.Contains(tag, ",") {
				sep := strings.SplitN(tag, ",", 2)
				tag = sep[0]
				split = strings.TrimSpace(sep[1])
				if split == "Space" {
					split = " "
				}
			}

			if j, ok := title[tag]; !ok {
				err = fmt.Errorf("%s is not in the title", tag)
				return
			} else {
				v := value.Field(i)

				var d string
				if len(row) > j {
					d = row[j]
				}

				if d == "" {
					continue
				}

				if split == "" {
					if err = cp(&v, d); err != nil {
						return
					}
				} else {
					vs := strings.Split(row[j], split)

					v.Set(reflect.MakeSlice(v.Type(), len(vs), len(vs)))
					for k, v1 := range vs {
						v.Index(k).SetString(strings.TrimSpace(v1))
					}
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
