package notice

import (
	"fmt"
	"os"
	"strings"
)

type MessageChain struct {
	stack []string
	value string
	bg    bool

	pipe *os.File
}

func (msg *MessageChain) getMode() int {
	if msg.bg {
		return modeBg
	}

	return modeText
}

func (msg *MessageChain) Pipe(pipe *os.File) {
	msg.pipe = pipe
}

func (msg *MessageChain) BG() *MessageChain {
	msg.bg = true
	return msg
}

func (msg *MessageChain) Text() *MessageChain {
	msg.bg = false
	return msg
}

func (msg *MessageChain) Get() string {
	msg.stack = append(msg.stack, msg.Reset().value)
	msg.value = ""

	value := strings.Join(msg.stack, "")
	return value
}

func (msg *MessageChain) Print() {
	if defaultPipe != nil {
		msg.pipe = defaultPipe
	}

	if msg.pipe == nil {
		msg.pipe = os.Stdout
	}

	_, _ = msg.pipe.WriteString(msg.Get())
	_, _ = msg.pipe.WriteString("\n")
}

func (msg *MessageChain) Panic() {
	panic(msg.Get())
}

func (msg *MessageChain) Message(value any) *MessageChain {
	msg.stack = append(msg.stack, msg.Reset().value)
	msg.value = fmt.Sprint(value)
	return msg
}

func (msg *MessageChain) Reset() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", reset.value, msg.value, reset.clear)
	return msg
}

func (msg *MessageChain) Bold() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", bold.value, msg.value, bold.clear)
	return msg
}

func (msg *MessageChain) Dim() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", dim.value, msg.value, dim.clear)
	return msg
}

func (msg *MessageChain) Italic() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", italic.value, msg.value, italic.clear)
	return msg
}

func (msg *MessageChain) Underline() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", underline.value, msg.value, underline.clear)
	return msg
}

func (msg *MessageChain) Inverse() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", inverse.value, msg.value, inverse.clear)
	return msg
}

func (msg *MessageChain) Hidden() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", hidden.value, msg.value, hidden.clear)
	return msg
}

func (msg *MessageChain) Strikethrough() *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", strikethrough.value, msg.value, strikethrough.clear)
	return msg
}

func (msg *MessageChain) RGB(r, g, b int) *MessageChain {
	mode := msg.getMode()
	msg.value = fmt.Sprintf("\u001b[%d;2;%d;%d;%dm%s\u001b[%dm", mode, r, g, b, msg.value, colorClear)
	return msg
}

func (msg *MessageChain) Hex(hex string) *MessageChain {
	mode := msg.getMode()
	r, g, b := hexToRGB(hex)
	msg.value = fmt.Sprintf("\u001b[%d;2;%d;%d;%dm%s\u001b[%dm", mode, r, g, b, msg.value, colorClear)
	return msg
}

func (msg *MessageChain) Hue(hue float64) *MessageChain {
	mode := msg.getMode()
	r, g, b := hueToRGB(hue)
	msg.value = fmt.Sprintf("\u001b[%d;2;%d;%d;%dm%s\u001b[%dm", mode, r, g, b, msg.value, colorClear)
	return msg
}

func (msg *MessageChain) HSL(hue, saturation, lightness int) *MessageChain {
	mode := msg.getMode()
	r, g, b := hslToRGB(hue, saturation, lightness)
	msg.value = fmt.Sprintf("\u001b[%d;2;%d;%d;%dm%s\u001b[%dm", mode, r, g, b, msg.value, colorClear)
	return msg
}

func (msg *MessageChain) color(color ansi) *MessageChain {
	msg.value = fmt.Sprintf("\u001b[%dm%s\u001b[%dm", color.value, msg.value, color.clear)
	return msg
}
