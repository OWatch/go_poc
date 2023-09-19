package error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestError(t *testing.T) {
	err := foo1()
	if err != nil {
		fmt.Printf("main test, err:%+v", err)
	}
}

func foo1() error {
	err := foo2()
	if err != nil {
		return errors.WithMessage(err, "foo2 err")
	}
	return nil
}

func foo2() error {
	return errors.New("foo1 err")
}
