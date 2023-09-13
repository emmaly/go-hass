package mediaPlayer

type DeviceClass string

const (
	DeviceClassNone     DeviceClass = ""
	DeviceClassTv       DeviceClass = "tv"
	DeviceClassSpeaker  DeviceClass = "speaker"
	DeviceClassReceiver DeviceClass = "receiver"
)

func (t DeviceClass) String() string {
	switch t {
	case DeviceClassNone:
		return ""
	case DeviceClassTv:
		return "TV"
	case DeviceClassSpeaker:
		return "Speaker"
	case DeviceClassReceiver:
		return "Receiver"
	default:
		return "Unknown"
	}
}
