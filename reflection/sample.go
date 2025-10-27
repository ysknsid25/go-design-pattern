package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// MarshalJSONLike は reflect を使って構造体を簡易的に JSON 文字列に変換する
func MarshalJSONLike(v interface{}) (string, error) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// 構造体以外はサポート外
	if rt.Kind() != reflect.Struct {
		return "", fmt.Errorf("unsupported type: %s", rt.Kind())
	}

	var sb strings.Builder
	sb.WriteString("{")

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		// フィールド名をJSONキーにする
		key := field.Name

		// jsonタグがあればそれを使う
		if tag := field.Tag.Get("json"); tag != "" {
			key = tag
		}

		// 値を文字列に変換
		var valStr string
		switch value.Kind() {
		case reflect.String:
			valStr = strconv.Quote(value.String()) // "で囲む
		case reflect.Int, reflect.Int64, reflect.Int32:
			valStr = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			valStr = fmt.Sprintf("%t", value.Bool())
		default:
			valStr = strconv.Quote(fmt.Sprintf("%v", value.Interface()))
		}

		// JSONフィールドを追加
		sb.WriteString(fmt.Sprintf(`"%s":%s`, key, valStr))

		if i < rt.NumField()-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString("}")
	return sb.String(), nil
}

type User struct {
	Name string `json:"name"` // Goの構造体にはタグをつけることができて、`field.Tag.Get("json");`みたいに判別できる
	Age  int    `json:"age"`
	VIP  bool   `json:"vip"`
}

func main() {
	u := User{Name: "Alice", Age: 30, VIP: true}

	jsonStr, err := MarshalJSONLike(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonStr)
}
