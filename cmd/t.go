package tt

import "github.com/wwq1988/errors"

func test() error {
	return errors.New("some")
}

func test1() error {
	err := test()
	return errors.Trace(err)
}
