package mediaPlayer

// RepeatMode is the Repeat Mode of the media, also known as repeat
type RepeatMode string

const (
	RepeatModeAll RepeatMode = "all"
	RepeatModeOff RepeatMode = "off"
	RepeatModeOne RepeatMode = "one"
)

func (t RepeatMode) String() string {
	switch t {
	case RepeatModeAll:
		return "All"
	case RepeatModeOff:
		return "Off"
	case RepeatModeOne:
		return "One"
	case "":
		return ""
	default:
		return "Unknown"
	}
}
