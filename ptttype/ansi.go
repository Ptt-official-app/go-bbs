package ptttype

func ANSIColor(color string) string {
	colorStr := ESC_STR + "[" + color + "m"
	return colorStr
}

func ANSIReset() string {
	return ESC_STR + "[m"
}
