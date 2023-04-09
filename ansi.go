package notice

type ansi struct {
	value int32
	clear int32
}

func mode(msg *MessageChain, textColor ansi, backgroundColor ansi) ansi {
	if msg.bg {
		return backgroundColor
	}

	return textColor
}

func (msg *MessageChain) Black() *MessageChain {
	return msg.color(mode(msg, black, bgBlack))
}

func (msg *MessageChain) Red() *MessageChain {
	return msg.color(mode(msg, red, bgRed))
}

func (msg *MessageChain) Green() *MessageChain {
	return msg.color(mode(msg, green, bgGreen))
}

func (msg *MessageChain) Yellow() *MessageChain {
	return msg.color(mode(msg, yellow, bgYellow))
}

func (msg *MessageChain) Blue() *MessageChain {
	return msg.color(mode(msg, blue, bgBlue))
}

func (msg *MessageChain) Magenta() *MessageChain {
	return msg.color(mode(msg, magenta, bgMagenta))
}

func (msg *MessageChain) Cyan() *MessageChain {
	return msg.color(mode(msg, cyan, bgCyan))
}

func (msg *MessageChain) White() *MessageChain {
	return msg.color(mode(msg, white, bgWhite))
}

func (msg *MessageChain) BlackBright() *MessageChain {
	return msg.color(mode(msg, blackBright, bgBlackBright))
}

func (msg *MessageChain) RedBright() *MessageChain {
	return msg.color(mode(msg, redBright, bgRedBright))
}

func (msg *MessageChain) GreenBright() *MessageChain {
	return msg.color(mode(msg, greenBright, bgGreenBright))
}

func (msg *MessageChain) YellowBright() *MessageChain {
	return msg.color(mode(msg, yellowBright, bgYellowBright))
}

func (msg *MessageChain) BlueBright() *MessageChain {
	return msg.color(mode(msg, blueBright, bgBlueBright))
}

func (msg *MessageChain) MagentaBright() *MessageChain {
	return msg.color(mode(msg, magentaBright, bgMagentaBright))
}

func (msg *MessageChain) CyanBright() *MessageChain {
	return msg.color(mode(msg, cyanBright, bgCyanBright))
}

func (msg *MessageChain) WhiteBright() *MessageChain {
	return msg.color(mode(msg, whiteBright, bgWhiteBright))
}
