package tool

import "fmt"

func Log(prefix string, content string) {
	fmt.Printf("%s.%s\n", prefix, content)
}
