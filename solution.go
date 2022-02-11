package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:")
}

func main() {
	fmt.Println(GetMessage())
}
