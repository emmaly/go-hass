package mediaPlayer

import (
	"encoding/json"
	"strings"

	"github.com/emmaly/go-hass"
)

type EntityPicture struct {
	Valid bool
	Value string
}

func (e *EntityPicture) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	e.Value = s
	e.Valid = true
	return nil
}

func (e EntityPicture) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return json.Marshal("")
	}
	return json.Marshal(e.Value)
}

func (e EntityPicture) String() string {
	if !e.Valid {
		return ""
	}
	return e.Value
}

func (e EntityPicture) URL(access *hass.Access) string {
	if !e.Valid {
		return ""
	}
	if strings.HasPrefix(e.Value, "/") {
		u, err := access.BuildURL("", "")
		if err != nil {
			return ""
		}
		return u + e.Value
	}
	return e.Value
}
