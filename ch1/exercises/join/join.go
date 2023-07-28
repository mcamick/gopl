package join

import (
	"fmt"
	"os"
	"strings"
)

func Join() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Println(os.Args[1:])
}
