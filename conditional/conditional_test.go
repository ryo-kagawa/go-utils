package conditional_test

import (
	"testing"

	"github.com/ryo-kagawa/go-utils/conditional"
)

func TestInt(t *testing.T) {
	trueValue := 1
	falseValue := 2
	if conditional.Int(true, trueValue, falseValue) != trueValue {
		t.Fatal("error to Int true")
	}
	if conditional.Int(false, trueValue, falseValue) != falseValue {
		t.Fatal("error to Int false")
	}
}

func TestString(t *testing.T) {
	trueValue := "true"
	falseValue := "false"
	if conditional.String(true, trueValue, falseValue) != trueValue {
		t.Fatal("error to String true")
	}
	if conditional.String(false, trueValue, falseValue) != falseValue {
		t.Fatal("error to String false")
	}
}
