package mediaPlayer

import (
	"encoding/json"
	"fmt"
	"math"
)

type MediaVolumeLevel struct {
	Value float64
	Valid bool
}

func (l *MediaVolumeLevel) UnmarshalJSON(b []byte) error {
	var s float64
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	l.Value = s
	l.Valid = true
	return nil
}

func (l MediaVolumeLevel) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(l.Value)
}

func (l *MediaVolumeLevel) String() string {
	if !l.Valid {
		return ""
	}
	return fmt.Sprintf("%0.0f%%", math.Round(l.Value*100))
}
