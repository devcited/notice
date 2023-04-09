package notice

import "fmt"

const clearLine = "\033[2K"

func ClearLine() {
	fmt.Printf(clearLine)
	fmt.Printf("\r")
}
