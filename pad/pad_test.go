package pad_test

import (
	"testing"

	"github.com/ryo-kagawa/go-utils/pad"
)

func TestLeft(t *testing.T) {
	value1 := "aaa"
	expected1 := "aaa"
	result1 := pad.Left(value1, "b", 1)
	if result1 != expected1 {
		t.Fatal("error to Left")
	}

	value2 := "aaa"
	expected2 := "bbbaaa"
	result2 := pad.Left(value2, "b", 6)
	if result2 != expected2 {
		t.Fatal("error to Left")
	}

	value3 := "aaa"
	expected3 := "bcbaaa"
	result3 := pad.Left(value3, "bc", 6)
	if result3 != expected3 {
		t.Fatal("error to Left")
	}
}

func TestRight(t *testing.T) {
	value1 := "aaa"
	expected1 := "aaa"
	result1 := pad.Right(value1, "b", 1)
	if result1 != expected1 {
		t.Fatal("error to Right")
	}

	value2 := "aaa"
	expected2 := "aaabbb"
	result2 := pad.Right(value2, "b", 6)
	if result2 != expected2 {
		t.Fatal("error to Right")
	}

	value3 := "aaa"
	expected3 := "aaabcb"
	result3 := pad.Right(value3, "bc", 6)
	if result3 != expected3 {
		t.Fatal("error to Right")
	}
}
