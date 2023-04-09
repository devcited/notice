package notice

import "os"

const (
	modeText = 38
	modeBg   = 48
)

var defaultPipe *os.File

var (
	reset         = ansi{0, 0}
	bold          = ansi{1, 22}
	dim           = ansi{2, 22}
	italic        = ansi{3, 23}
	underline     = ansi{4, 24}
	inverse       = ansi{7, 27}
	hidden        = ansi{8, 28}
	strikethrough = ansi{9, 29}
)

var (
	colorClear   int32 = 39
	bgColorClear int32 = 49
)

var (
	black   = ansi{30, colorClear}
	red     = ansi{31, colorClear}
	green   = ansi{32, colorClear}
	yellow  = ansi{33, colorClear}
	blue    = ansi{34, colorClear}
	magenta = ansi{35, colorClear}
	cyan    = ansi{36, colorClear}
	white   = ansi{37, colorClear}

	blackBright   = ansi{90, colorClear}
	redBright     = ansi{91, colorClear}
	greenBright   = ansi{92, colorClear}
	yellowBright  = ansi{93, colorClear}
	blueBright    = ansi{94, colorClear}
	magentaBright = ansi{95, colorClear}
	cyanBright    = ansi{96, colorClear}
	whiteBright   = ansi{97, colorClear}
)

var (
	bgBlack   = ansi{40, bgColorClear}
	bgRed     = ansi{41, bgColorClear}
	bgGreen   = ansi{42, bgColorClear}
	bgYellow  = ansi{43, bgColorClear}
	bgBlue    = ansi{44, bgColorClear}
	bgMagenta = ansi{45, bgColorClear}
	bgCyan    = ansi{46, bgColorClear}
	bgWhite   = ansi{47, bgColorClear}

	bgBlackBright   = ansi{100, bgColorClear}
	bgRedBright     = ansi{101, bgColorClear}
	bgGreenBright   = ansi{102, bgColorClear}
	bgYellowBright  = ansi{103, bgColorClear}
	bgBlueBright    = ansi{104, bgColorClear}
	bgMagentaBright = ansi{105, bgColorClear}
	bgCyanBright    = ansi{106, bgColorClear}
	bgWhiteBright   = ansi{107, bgColorClear}
)
