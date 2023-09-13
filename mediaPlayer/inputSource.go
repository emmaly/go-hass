package mediaPlayer

import (
	"encoding/json"
)

type InputSourceList struct {
	Valid bool
	Value []InputSource
}

func (l *InputSourceList) UnmarshalJSON(b []byte) error {
	var s []InputSource
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	l.Value = s
	l.Valid = true
	return nil
}

func (l InputSourceList) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return json.Marshal("")
	}
	return json.Marshal(l.Value)
}

func (l *InputSourceList) String() string {
	if !l.Valid {
		return ""
	}
	if len(l.Value) == 0 {
		return "(none)"
	}
	var s string
	for _, v := range l.Value {
		s += string(v) + " "
	}
	return s
}

type InputSource string
