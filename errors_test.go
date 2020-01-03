package errors

import (
	"errors"

	"github.com/wwq1988/errors/stack"
)

var (
	err1 = errors.New("err1")
	err2 = errors.New("err2")
	err3 = errors.New("err3")
)

func innerTest1(err error) error {
	return NewEx(1, err, nil)
}

func innerTest2(err error) error {
	return NewEx(1, err, nil)
}

func innerTest3(err error) error {
	return NewEx(1, err, nil)
}

func traceTest1(err error) error {
	return Trace(innerTest1(err))
}

func traceTest2(err error) error {
	return Trace(innerTest2(err))
}

func traceTest3(err error) error {
	return Trace(innerTest3(err))
}

func traceWithFieldTest1(err error, key string, value interface{}) error {
	return TraceWithField(innerTest1(err), key, value)
}

func traceWithFieldTest2(err error, key string, value interface{}) error {
	return TraceWithField(innerTest2(err), key, value)
}

func traceWithFieldTest3(err error, key string, value interface{}) error {
	return TraceWithField(innerTest3(err), key, value)
}

func traceWithFieldsTest1(err error, fs stack.Fields) error {
	return TraceWithFields(innerTest1(err), fs)
}

func traceWithFieldsTest2(err error, fs stack.Fields) error {
	return TraceWithFields(innerTest2(err), fs)
}

func traceWithFieldsTest3(err error, fs stack.Fields) error {
	return TraceWithFields(innerTest3(err), fs)
}

// func TestInner(t *testing.T) {
// 	type Field struct {
// 		Key   string
// 		Value interface{}
// 	}
// 	tests := []struct {
// 		rawErr        error
// 		fn            func(error) error
// 		expectedStack string
// 		givenFields   stack.Fields
// 		givenField    Field
// 	}{
// 		{rawErr: err1, fn: innerTest1, expectedStack: "errors.innerTest1:18", givenField: Field{Key: "key1", Value: "key1"}, givenFields: stack.FromKVs(map[string]interface{}{"key11": "value11", "key12": "value12"})},
// 		{rawErr: err2, fn: innerTest2, expectedStack: "errors.innerTest2:22", givenField: Field{Key: "key2", Value: "key2"}, givenFields: stack.FromKVs(map[string]interface{}{"key21": "value21", "key22": "value22"})},
// 		{rawErr: err3, fn: innerTest3, expectedStack: "errors.innerTest3:26", givenField: Field{Key: "key3", Value: "key3"}, givenFields: stack.FromKVs(map[string]interface{}{"key31": "value31", "key32": "value32"})},
// 	}

// 	for _, test := range tests {
// 		err := test.fn(test.rawErr)
// 		stackErr := err.(*StackError)
// 		gotStack := stackErr.Fields()
// 		if gotStack != test.expectedStack {
// 			t.Fatalf("given:%+v,expectedstack:%s,gotstack:%s", err, test.expectedStack, gotStack)
// 		}
// 		gotFields := stackErr.Fields().KVs()
// 		expectedFields := map[string]interface{}{"stack": test.expectedStack}
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}
// 		stackErr.cleanFields()
// 		stackErr.fillStackField()
// 		gotFields = stackErr.fields.KVs()
// 		expectedFields = map[string]interface{}{"stack": test.expectedStack}
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}
// 		stackErr.cleanFields()
// 		gotFields = stackErr.Fields().KVs()
// 		expectedFields = map[string]interface{}{}
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}
// 		stackErr.cleanFields()
// 		stackErr.WithField(test.givenField.Key, test.givenField.Value)
// 		expectedFields = map[string]interface{}{test.givenField.Key: test.givenField.Value}
// 		gotFields = stackErr.Fields().KVs()
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}
// 		stackErr.cleanFields()
// 		stackErr.WithFields(test.givenFields)
// 		gotFields = stackErr.Fields().KVs()
// 		expectedFields = test.givenFields.KVs()
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}
// 		stackErr.cleanFields()
// 		gotErr := stackErr.Unwrap()
// 		if stackErr.Unwrap() != test.rawErr {
// 			t.Fatalf("given:%+v,expectederr:%+v,goterr:%+v", err, test.rawErr, gotErr)
// 		}
// 		if !stackErr.Is(test.rawErr) {
// 			t.Fatalf("given:%+v,expectedis:true,gotis:false", err)
// 		}
// 		gotErrStr := stackErr.Error()
// 		expectedErrStr := test.rawErr.Error()
// 		if gotErrStr != expectedErrStr {
// 			t.Fatalf("given:%+v,expectederrstr:%+v,goterrstr:%+v", err, expectedErrStr, gotErrStr)
// 		}
// 	}
// }

// func TestTrace(t *testing.T) {
// 	tests := []struct {
// 		rawErr        error
// 		fn            func(error) error
// 		expectedStack string
// 	}{
// 		{rawErr: err1, fn: traceTest1, expectedStack: "errors.innerTest1:18;errors.traceTest1:30"},
// 		{rawErr: err2, fn: traceTest2, expectedStack: "errors.innerTest2:22;errors.traceTest2:34"},
// 		{rawErr: err3, fn: traceTest3, expectedStack: "errors.innerTest3:26;errors.traceTest3:38"},
// 	}
// 	for _, test := range tests {
// 		err := test.fn(test.rawErr)
// 		stackErr := err.(*StackError)
// 		gotStack := stackErr.stack()
// 		if gotStack != test.expectedStack {
// 			t.Fatalf("given:%+v,expectedstack:%s,gotstack:%s", err, test.expectedStack, gotStack)
// 		}
// 		gotFields := stackErr.Fields().KVs()
// 		expectedFields := map[string]interface{}{"stack": test.expectedStack}
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}

// 	}
// }

// func TestTraceWithField(t *testing.T) {
// 	type Field struct {
// 		Key   string
// 		Value interface{}
// 	}
// 	tests := []struct {
// 		rawErr        error
// 		fn            func(error, string, interface{}) error
// 		expectedStack string
// 		givenField    Field
// 	}{
// 		{rawErr: err1, fn: traceWithFieldTest1, expectedStack: "errors.innerTest1:18;errors.traceWithFieldTest1:42", givenField: Field{Key: "key1", Value: "key1"}},
// 		{rawErr: err2, fn: traceWithFieldTest2, expectedStack: "errors.innerTest2:22;errors.traceWithFieldTest2:46", givenField: Field{Key: "key2", Value: "key2"}},
// 		{rawErr: err3, fn: traceWithFieldTest3, expectedStack: "errors.innerTest3:26;errors.traceWithFieldTest3:50", givenField: Field{Key: "key3", Value: "key3"}},
// 	}
// 	for _, test := range tests {
// 		err := test.fn(test.rawErr, test.givenField.Key, test.givenField.Value)
// 		stackErr := err.(*StackError)
// 		gotStack := stackErr.stack()
// 		if gotStack != test.expectedStack {
// 			t.Fatalf("given:%+v,expectedstack:%s,gotstack:%s", err, test.expectedStack, gotStack)
// 		}
// 		stackErr.cleanFields()
// 		gotFields := stackErr.Fields().KVs()
// 		expectedFields := map[string]interface{}{"stack": test.expectedStack}
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}

// 	}
// }

// func TestTraceWithFields(t *testing.T) {
// 	type Field struct {
// 		Key   string
// 		Value interface{}
// 	}
// 	tests := []struct {
// 		rawErr        error
// 		fn            func(error, stack.Fields) error
// 		expectedStack string
// 		givenFields   stack.Fields
// 	}{
// 		{rawErr: err1, fn: traceWithFieldsTest1, expectedStack: "errors.innerTest1:18;errors.traceWithFieldsTest1:54", givenFields: stack.FromKVs(map[string]interface{}{"key11": "value11", "key12": "value12"})},
// 		{rawErr: err2, fn: traceWithFieldsTest2, expectedStack: "errors.innerTest2:22;errors.traceWithFieldsTest2:58", givenFields: stack.FromKVs(map[string]interface{}{"key21": "value21", "key22": "value22"})},
// 		{rawErr: err3, fn: traceWithFieldsTest3, expectedStack: "errors.innerTest3:26;errors.traceWithFieldsTest3:62", givenFields: stack.FromKVs(map[string]interface{}{"key31": "value31", "key22": "value32"})},
// 	}
// 	for _, test := range tests {
// 		err := test.fn(test.rawErr, test.givenFields)
// 		stackErr := err.(*StackError)
// 		gotStack := stackErr.stack()
// 		if gotStack != test.expectedStack {
// 			t.Fatalf("given:%+v,expectedstack:%s,gotstack:%s", err, test.expectedStack, gotStack)
// 		}
// 		stackErr.cleanFields()
// 		gotFields := stackErr.Fields().KVs()
// 		expectedFields := map[string]interface{}{"stack": test.expectedStack}
// 		if !reflect.DeepEqual(gotFields, expectedFields) {
// 			t.Fatalf("given:%+v,expectedfields:%+v,gotfields:%+v", err, expectedFields, gotFields)
// 		}
// 	}
// }

// func TestUnwrap(t *testing.T) {
// 	tests := []struct {
// 		rawErr error
// 		fn     func(error) error
// 	}{
// 		{rawErr: err1, fn: innerTest1},
// 		{rawErr: err2, fn: innerTest2},
// 		{rawErr: err3, fn: innerTest3},
// 	}
// 	for _, test := range tests {
// 		err := test.fn(test.rawErr)
// 		gotErr := Unwrap(err)
// 		if Unwrap(err) != test.rawErr {
// 			t.Fatalf("given:%+v,expectederr:%+v,goterr:%+v", err, test.rawErr, gotErr)
// 		}
// 	}
// }

// func TestIs(t *testing.T) {
// 	tests := []struct {
// 		rawErr error
// 		fn     func(error) error
// 	}{
// 		{rawErr: err1, fn: innerTest1},
// 		{rawErr: err2, fn: innerTest2},
// 		{rawErr: err3, fn: innerTest3},
// 	}
// 	for _, test := range tests {
// 		err := test.fn(test.rawErr)
// 		if !Is(err, test.rawErr) {
// 			t.Fatalf("given:%+v,expectedis:true,gotis:false", err)
// 		}
// 	}
// }
