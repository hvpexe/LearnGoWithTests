package strings

import (
	"strings"
	"testing"
	"unsafe"
)

func TestClone(t *testing.T) {
	t.Run("check origin and clone must be the same", func(t *testing.T) {
		origin := "aaa"
		clone := strings.Clone(origin)
		if origin != clone {
			t.Errorf("origin != clone")
		}
	})
	t.Run("check origin and clone must be in different memory", func(t *testing.T) {
		origin := "aaa"
		clone := strings.Clone(origin)
		if unsafe.StringData(origin) == unsafe.StringData(clone) {
			t.Errorf("origin and clone are in the same memory")
		}
	})
}
