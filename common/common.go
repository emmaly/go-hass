package common

type Common struct {
	EntityID   string `json:"entity_id,omitempty"`
	EntityArea string `json:"entity_area,omitempty"`
	DeviceID   string `json:"device_id,omitempty"`
	DeviceArea string `json:"device_area,omitempty"`
	Name       string `json:"name,omitempty"`
	State      string `json:"state,omitempty"`
}

type CommonAttributes struct {
	DeviceClass  string `json:"device_class,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	IPAddress    string `json:"ip_address,omitempty"`
}
