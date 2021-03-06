package e

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ErrClassInternal = iota
	ErrClassUser
)

func TestSimpleError(t *testing.T) {
	e := NewError(ErrClassInternal, "a")
	assert.EqualError(t, e, "a")
}

func TestSimpleErrorWithFormat(t *testing.T) {
	e := NewErrorf(ErrClassInternal, "a%s", "b")
	assert.EqualError(t, e, "ab")
}

func TestWrappedErrorWithMessage(t *testing.T) {
	i := errors.New("a")
	e := Wrapf(ErrClassInternal, i, "b")
	assert.EqualError(t, e, "b")
	assert.EqualError(t, e.InnerError(), "a")
}

func TestWrappedError(t *testing.T) {
	e := Wrap(ErrClassInternal, errors.New("a"))
	assert.EqualError(t, e, "a")
}

func TestStack(t *testing.T) {
	e := Wrap(ErrClassInternal, errors.New("b"))
	assert.Equal(t, "github.com/mbtproject/mbt/e.TestStack", e.Stack()[1].Function)
}

func TestExtendedInfo(t *testing.T) {
	err := NewError(ErrClassInternal, "blah")
	assert.Contains(t, err.WithExtendedInfo().Error(), "call stack")
}

func WrappingAnE(t *testing.T) {
	a := Wrap(ErrClassInternal, errors.New("a"))
	assert.Equal(t, a, Wrap(ErrClassInternal, a))
	assert.Equal(t, a, Wrap(ErrClassUser, a))
}
