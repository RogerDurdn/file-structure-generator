package pkg

import (
	"reflect"
	"testing"
)

func Test_obtainRecordsFromFile(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		name := obtainRecordsFromFile("/dataTab.csv")[0].Name
		if got := name; !reflect.DeepEqual(got, "a1") {
			t.Errorf("obtainRecordsFromFile() = %v, want %v", got, "a1")
		}
	})
}

func TestGenerateTree(t *testing.T) {
	t.Run("testGenerateTree", func(t *testing.T) {
		Generate("/dataTab.csv")
	})
}
