package dir

import (
	"testing"
	//	"github.com/aizu-go-kapro/dir"
)

func TestDo(t *testing.T) {
	cases := []struct {
		name string
		path string
	}{
		{"TestLocal", "./testcases/test1.sh"},
		{"TestRemote", "./testcases/test2.sh"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			executor := NewExecutor(c.path)
			if err := executor.Do(); err != nil {
				t.Errorf("Failed to execute %s ", c.path)
			}
		})
	}
}
