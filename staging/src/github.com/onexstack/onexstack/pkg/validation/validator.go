package validation

import (
	"fmt"
	"reflect"
)

// 定义验证函数类型
type ValidatorFunc func(value any) error

// 定义验证规则类型
type Rules map[string]ValidatorFunc

func ValidateAllFields(obj any, rules Rules) error {
	return ValidateSelectedFields(obj, rules, GetExportedFieldNames(obj)...)
}

// 通用校验函数
func ValidateSelectedFields(obj any, rules Rules, fields ...string) error {
	// 通过反射获取结构体的值和类型
	objValue := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	// 确保传入的是一个结构体
	if objType.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}

	// 确保传入的是一个结构体
	if objType.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %s", objType.Kind())
	}

	// 遍历需要校验的字段
	for _, field := range fields {
		// 检查字段是否存在
		structField, exists := objType.FieldByName(field)
		if !exists || !structField.IsExported() {
			continue // 跳过不存在或未导出的字段
		}

		// 提取字段值
		fieldValue := objValue.FieldByName(field)

		// 校验字段值是否是指针
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				// 如果是指针类型但为 nil，返回错误
				// return fmt.Errorf("field '%s' cannot be nil", field)
				// 如果是指针类型但为 nil，跳过验证
				continue
			}
			// 如果是指针，取消引用，获取指针的实际值
			fieldValue = fieldValue.Elem()
		}

		// 获取字段的实际值并校验
		validator, ok := rules[field]
		if !ok {
			continue // 跳过没有校验规则的字段
		}

		if err := validator(fieldValue.Interface()); err != nil {
			return err
		}
	}

	return nil
}

// GetExportedFieldNames 返回传入结构体中所有可导出的字段名字.
func GetExportedFieldNames(obj any) []string {
	// 获取传入值的类型和值
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	// 校验是否为结构体或结构体指针
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return []string{}
	}

	var fieldNames []string
	// 遍历结构体的字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 只添加可导出字段（字段名以大写字母开头）
		if field.IsExported() { // 从 Go 1.17 开始提供
			fieldNames = append(fieldNames, field.Name)
		}
	}

	return fieldNames
}
