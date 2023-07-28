package util

import (
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// CheckState 校验状态是否有效
func CheckState(state int) bool {
	return state == StateEnable || state == StateDisable || state == StateDelete
}

// CheckObjType 校验对象类型是否有效
func CheckObjType(objType string) bool {
	return objType == ObjTypeCate || objType == ObjTypeArticle || objType == ObjTypeUser
}

// GenRandomString 生成指定长度的随机字符串
func GenRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	source := "0123456789abcdefghijklmnopqrstuvwxyz"
	sourceLen := len(source)
	for i := 0; i < length; i++ {
		randIdx := rand.Intn(sourceLen)
		result[i] = source[randIdx]
	}

	return string(result)
}

// Struct2Struct 结构体转结构体
func Struct2Struct(source, target any) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	// 校验源、目标数据类型
	if sourceValue.Kind() != reflect.Struct || targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Struct {
		return
	}

	sourceType := reflect.TypeOf(source)
	targetType := reflect.TypeOf(target).Elem()

	// 目标数据字段名索引
	fieldNum := targetType.NumField()
	targetFieldMap := make(map[string]int, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := targetType.Field(i).Tag.Get("json")
		if tagName == "" {
			continue
		}

		tagName = strings.Split(tagName, ",")[0]
		targetFieldMap[tagName] = i
	}

	// 遍历源数据字段
	fieldNum = reflect.TypeOf(source).NumField()
	for i := 0; i < fieldNum; i++ {
		tagName := sourceType.Field(i).Tag.Get("json")
		if tagName == "" {
			continue
		}
		tagName = strings.Split(tagName, ",")[0]

		// 校验目标数据字段是否存在
		targetFieldIndex, ok := targetFieldMap[tagName]
		if !ok {
			continue
		}

		sourceField := sourceValue.Field(i)

		// 校验目标数据字段类型是否与源数据字段类型一致
		targetField := targetValue.Elem().Field(targetFieldIndex)
		if targetField.Kind() != sourceField.Kind() {
			continue
		}

		targetField.Set(sourceField)
	}
}
