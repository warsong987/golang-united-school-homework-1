package solution_test

import (
	"golang-united-school-homework-1/solution"
	"testing"

	"github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	res := solution.GetMessage()
	wait := emoji.Sprint("Hello :world_map:")

	if res != wait {
		t.Fatalf(`Message does "%s" not match the message "%s"`, res, wait)
	}
}
