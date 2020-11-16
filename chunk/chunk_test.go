package chunk_test

import (
	"reflect"
	"testing"

	"github.com/ryo-kagawa/go-utils/chunk"
)

func TestInt(t *testing.T) {
	value := []int{0, 1, 2, 3, 4}
	expected := [][]int{{0, 1}, {2, 3}, {4}}
	result := chunk.Int(value, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("error to Int")
	}
}
