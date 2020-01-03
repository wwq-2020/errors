package errors

import (
	"errors"
	"fmt"

	"github.com/wwq1988/errors/stack"
)

// timeout 判断是否超时接口
type timeout interface {
	Timeout() bool
}

// temporary 判断是否临时接口
type temporary interface {
	Temporary() bool
}

// StackError StackError
type StackError struct {
	fields map[string]stack.Fields
	err    error
}

// New 初始化StackError
func New(msg string, args ...interface{}) error {
	return NewEx(2, fmt.Errorf(msg, args...), nil)
}

// NewWithFields 初始化StackError带域存储
func NewWithFields(msg string, fields stack.Fields) error {
	return NewEx(2, fmt.Errorf(msg), fields)
}

// NewWithField 初始化StackError带kv
func NewWithField(msg string, key string, val interface{}) error {
	fs := stack.New().Set(key, val)
	return NewEx(2, fmt.Errorf(msg), fs)
}

// NewEx 初始化StackError,带堆栈深度和域存储
func NewEx(depth int, err error, fields stack.Fields) error {
	stackFrame := stack.Get(depth)
	if fields == nil {
		fields = stack.New()
	}
	se, ok := err.(*StackError)
	if !ok {
		se = &StackError{
			err:    err,
			fields: make(map[string]stack.Fields),
		}
	}
	se.fields[stackFrame] = fields
	return se
}

// Fields 获取域存储
func (s *StackError) Fields() map[string]stack.Fields {
	return s.fields
}

// Unwrap Unwrap
func (s *StackError) Unwrap() error {
	return s.err
}

// Is Is
func (s *StackError) Is(err error) bool {
	se, ok := err.(*StackError)
	return ok && se.err == s.err
}

// As As
func (s *StackError) As(err error) bool {
	se, ok := err.(*StackError)
	if ok {
		*se = *s
	}
	return ok
}

// Error Error
func (s *StackError) Error() string {
	return s.err.Error()
}

// Trace 追踪错误
func Trace(err error) error {
	if err == nil {
		return nil
	}
	return NewEx(2, err, nil)
}

// TraceWithFields 追踪错误带域存储
func TraceWithFields(err error, fields stack.Fields) error {
	return TraceWithFieldsEx(err, fields, 2)
}

// TraceWithFieldsEx 追踪错误带域存储和堆栈深度
func TraceWithFieldsEx(err error, fields stack.Fields, depth int) error {
	if err == nil {
		return nil
	}
	return NewEx(depth+1, err, fields)
}

// TraceWithFieldEx 追踪错误带kv和堆栈深度
func TraceWithFieldEx(err error, key string, val interface{}, depth int) error {
	fs := stack.New()
	fs.Set(key, val)
	return TraceWithFieldsEx(err, fs, depth+1)
}

// TraceWithField 追踪错误带kv
func TraceWithField(err error, key string, val interface{}) error {
	return TraceWithFieldEx(err, key, val, 2)

}

// Is 判断err是否相同
func Is(src, dst error) bool {
	return errors.Is(src, dst)
}

// Unwrap 返回原始err
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// IsTimeout 判断是否超时
func IsTimeout(err error) bool {
	timeoutErr, ok := err.(timeout)
	return ok && timeoutErr.Timeout()
}

// IsTemporary 判断是否是临时错误
func IsTemporary(err error) bool {
	temporaryErr, ok := err.(temporary)
	return ok && temporaryErr.Temporary()
}

// Fields 获取域
func Fields(err error) map[string]stack.Fields {
	if stackErr, ok := err.(*StackError); ok {
		return stackErr.Fields()
	}
	return nil
}
