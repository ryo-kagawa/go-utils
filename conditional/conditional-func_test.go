package conditional_test

import (
	"testing"

	"github.com/ryo-kagawa/go-utils/conditional"
)

func TestIntFunc(t *testing.T) {
	trueValue := 1
	falseValue := 2
	trueFunc := func() int { return trueValue }
	falseFunc := func() int { return falseValue }
	if conditional.IntFunc(true, trueFunc, falseFunc) != trueValue {
		t.Fatal("error to Int true")
	}
	if conditional.IntFunc(false, trueFunc, falseFunc) != falseValue {
		t.Fatal("error to Int false")
	}
}

func TestStringFunc(t *testing.T) {
	trueValue := "true"
	falseValue := "false"
	trueFunc := func() string { return trueValue }
	falseFunc := func() string { return falseValue }
	if conditional.StringFunc(true, trueFunc, falseFunc) != trueValue {
		t.Fatal("error to String true")
	}
	if conditional.StringFunc(false, trueFunc, falseFunc) != falseValue {
		t.Fatal("error to String false")
	}
}
