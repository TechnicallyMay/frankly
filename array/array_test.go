package array
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseArrayReversesCorrectly(t *testing.T) {
    toReverse := [...]string{"1", "2", "3"}
    toReverseCopy := [...]string{"1", "2", "3"}

    ReverseArray(toReverse[:])

    for i, j := 0, len(toReverseCopy) - 1; j >=0; i, j = i+1, j-1 {
        initialValue := toReverseCopy[i]
        reversedValue := toReverse[j]
        assert.Equal(t, initialValue, reversedValue, "Initial list at index [%v] was %v. Expected reversed list at index [%v] to match, was %v", i, initialValue, j, reversedValue)
    }
}

