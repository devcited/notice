package notice

import (
	"fmt"
	"os"
)

func Pipe(pipe *os.File) {
	defaultPipe = pipe
}

func Message(value any) *MessageChain {
	msg := MessageChain{
		[]string{},
		fmt.Sprint(value),
		false,
		os.Stdout,
	}

	return &msg
}
