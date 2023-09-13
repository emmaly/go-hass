package mediaPlayer

import "encoding/json"

type MediaVolumeMuted struct {
	Value bool
	Valid bool
}

func (l *MediaVolumeMuted) UnmarshalJSON(b []byte) error {
	var s bool
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	l.Value = s
	l.Valid = true
	return nil
}

func (l MediaVolumeMuted) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(l.Value)
}

func (l *MediaVolumeMuted) String() string {
	if !l.Valid {
		return ""
	}
	if l.Value {
		return "muted"
	}
	return "unmuted"
}
