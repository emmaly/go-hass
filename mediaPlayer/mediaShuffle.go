package mediaPlayer

import "encoding/json"

// MediaShuffle is the Shuffle Mode of the media, also known as shuffle
type MediaShuffle struct {
	Value bool
	Valid bool
}

func (l *MediaShuffle) UnmarshalJSON(b []byte) error {
	var s bool
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	l.Value = s
	l.Valid = true
	return nil
}

func (l MediaShuffle) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(l.Value)
}

func (l *MediaShuffle) String() string {
	if !l.Valid {
		return ""
	}
	if l.Value {
		return "shuffle on"
	}
	return "shuffle off"
}
