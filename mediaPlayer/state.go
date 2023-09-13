package mediaPlayer

type State string

const (
	StateOff       State = "off"
	StateOn        State = "on"
	StateIdle      State = "idle"
	StatePlaying   State = "playing"
	StatePaused    State = "paused"
	StateStandby   State = "standby"
	StateBuffering State = "buffering"
)
