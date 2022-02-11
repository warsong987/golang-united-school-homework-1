package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	res := GetMessage()
	wait := emoji.Sprint("Hello :world_map:")

	if res != wait {
		t.Fatalf(`Message does "%s" not match the message "%s"`, res, wait)
	}
}
