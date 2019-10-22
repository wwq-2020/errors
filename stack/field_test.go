package stack

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	tests := []struct {
		Key   string
		Value interface{}
	}{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
		{Key: "key3", Value: "value3"},
	}
	for _, test := range tests {
		fs := New()
		fs.Set(test.Key, test.Value)
		expected := map[string]interface{}{test.Key: test.Value}
		got := (map[string]interface{})(fs.(fields))
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("given:%s,%+v,expected:%+v,got:%+v", test.Key, test.Value, expected, got)
		}
	}
}

func TestKVs(t *testing.T) {
	tests := []struct {
		expected map[string]interface{}
	}{
		{expected: map[string]interface{}{"key11": "value11"}},
		{expected: map[string]interface{}{"key21": "value21"}},
		{expected: map[string]interface{}{"key31": "value31"}},
	}
	for _, test := range tests {
		fs := fields(test.expected)

		got := fs.KVs()
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("given:%+v,expected:%+v,got:%+v", test.expected, test.expected, got)
		}
	}
}

func TestMerge(t *testing.T) {
	type Field struct {
		Key   string
		Value interface{}
	}
	tests := []struct {
		Dst Field
		Src Field
	}{
		{Dst: Field{"key11", "value11"}, Src: Field{"key12", "value12"}},
		{Dst: Field{"key21", "value21"}, Src: Field{"key22", "value22"}},
		{Dst: Field{"key31", "value31"}, Src: Field{"key32", "value32"}},
	}
	for _, test := range tests {
		dst := New().Set(test.Dst.Key, test.Dst.Value)
		src := New().Set(test.Src.Key, test.Src.Value)
		dst.Merge(src)
		expected := map[string]interface{}{test.Dst.Key: test.Dst.Value, test.Src.Key: test.Src.Value}
		got := (map[string]interface{})(dst.(fields))
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("given:%s,%+v,%s,%+v,expected:%+v,got:%+v", test.Dst.Key, test.Dst.Value, test.Src.Key, test.Src.Value, expected, got)
		}
	}
}

func TestFromKVs(t *testing.T) {

	tests := []struct {
		expected map[string]interface{}
	}{
		{expected: map[string]interface{}{"key11": "value11"}},
		{expected: map[string]interface{}{"key21": "value21"}},
		{expected: map[string]interface{}{"key31": "value31"}},
	}
	for _, test := range tests {
		fs := FromKVs(test.expected)
		got := (map[string]interface{})(fs.(fields))
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("given:%+v,expected:%+v,got:%+v", test.expected, test.expected, got)
		}
	}
}
