package common

import "encoding/json"

type Service struct {
	ServiceID         string
	Name              string
	Description       string
	Target            []TargetSelector
	Fields            []ServiceField
	SupportedFeatures []SupportedFeature
}

type ServiceField struct {
	FieldID     string
	Name        string
	Description string
	Required    bool
	Example     interface{}
	Selector    []Selector
	Filter      *ServiceFieldFilter
}

type SupportedFeature string

//

type Selector struct {
	Type              SelectorType
	Area              *AreaSelector
	Attribute         *AttributeSelector
	ColorTemp         *ColorTempSelector
	ConfigEntry       *ConfigEntrySelector
	Constant          *ConstantSelector
	ConversationAgent *ConversationAgentSelector
	Device            *DeviceSelector
	Duration          *DurationSelector
	Entity            *EntitySelector
	Icon              *IconSelector
	Language          *LanguageSelector
	Location          *LocationSelector
	Number            *NumberSelector
	Select            *SelectSelector
	State             *SelectSelector
	Target            *TargetSelector
	Text              *TextSelector
	Theme             *ThemeSelector
}

//

type SelectorType string

const (
	SelectorTypeAction            SelectorType = "action" // no options
	SelectorTypeArea              SelectorType = "area"
	SelectorTypeAttribute         SelectorType = "attribute"
	SelectorTypeAssistPipeline    SelectorType = "assist_pipeline" // no options
	SelectorTypeBool              SelectorType = "bool"            // no options
	SelectorTypeColorRGB          SelectorType = "color_rgb"       // no options
	SelectorTypeColorTemp         SelectorType = "color_temp"
	SelectorTypeCondition         SelectorType = "condition" // no options
	SelectorTypeConfigEntry       SelectorType = "config_entry"
	SelectorTypeConstant          SelectorType = "constant"
	SelectorTypeConversationAgent SelectorType = "conversation_agent"
	SelectorTypeDate              SelectorType = "date"     // no options
	SelectorTypeDateTime          SelectorType = "datetime" // no options
	SelectorTypeDevice            SelectorType = "device"
	SelectorTypeDuration          SelectorType = "duration"
	SelectorTypeEntity            SelectorType = "entity"
	SelectorTypeIcon              SelectorType = "icon"
	SelectorTypeLanguage          SelectorType = "language"
	SelectorTypeLocation          SelectorType = "location"
	SelectorTypeMedia             SelectorType = "media" // no options
	SelectorTypeNumber            SelectorType = "number"
	SelectorTypeObject            SelectorType = "object" // no options
	SelectorTypeSelect            SelectorType = "select"
	SelectorTypeState             SelectorType = "state"
	SelectorTypeTarget            SelectorType = "target"
	SelectorTypeTemplate          SelectorType = "template" // no options
	SelectorTypeText              SelectorType = "text"
	SelectorTypeTheme             SelectorType = "theme"
	SelectorTypeTime              SelectorType = "time" // no options
)

//

type AreaSelector struct {
	Device   []SelectorDeviceFilter
	Entity   []SelectorEntityFilter
	Multiple bool
}

type AttributeSelector struct {
	EntityID string `json:"entity_id,omitempty"`
}

type ColorTempSelector struct {
	MinMireds int `json:"min_mireds,omitempty"`
	MaxMireds int `json:"max_mireds,omitempty"`
}

type ConfigEntrySelector struct {
	Integration string
}

type ConstantSelector struct {
	Value bool
	Label string
}

type ConversationAgentSelector struct {
	Language string
}

type DeviceSelector struct {
	Filter   []SelectorDeviceFilter
	Entity   []SelectorEntityFilter
	Multiple bool
}

type DurationSelector struct {
	EnableDay bool `json:"enable_day"`
}

type EntitySelector struct {
	ExcludeEntities []string // exclude from the resulting list
	IncludeEntities []string // constrain to this list
	Filter          []SelectorEntityFilter
	Multiple        bool
}

type IconSelector struct {
	Placeholder string `json:"placeholder,omitempty"`
}

type LanguageSelector struct {
	Languages  []string `json:"languages,omitempty"` // RFC 5646 language codes
	NativeName bool     `json:"native_name"`
	NoSort     bool     `json:"no_sort"`
}

type LocationSelector struct {
	Icon   string `json:"icon,omitempty"`
	Radius bool   `json:"radius"`
}

type NumberSelector struct {
	Min               float64            `json:"min"`
	Max               float64            `json:"max"`
	Step              float64            `json:"step,omitempty"`
	UnitOfMeasurement string             `json:"unit_of_measurement,omitempty"`
	Mode              NumberSelectorMode `json:"mode,omitempty"`
}

type NumberSelectorMode string

const (
	NumberSelectorModeSlider NumberSelectorMode = "slider"
	NumberSelectorModeBox    NumberSelectorMode = "box"
)

type SelectSelector struct {
	Options        []string           `json:"options,omitempty"`
	OptionsKV      map[string]string  `json:"-"` // if set, this takes precedence over Options when marshaling
	Multiple       bool               `json:"multiple"`
	CustomValue    bool               `json:"custom_value"`
	Mode           SelectSelectorMode `json:"mode,omitempty"`
	TranslationKey string             `json:"translation_key,omitempty"`
	Sort           bool               `json:"sort"`
}

func (s *SelectSelector) MarshalJSON() ([]byte, error) {
	if s.OptionsKV != nil {
		type KV struct {
			Label string `json:"label"`
			Value string `json:"value"`
		}
		kv := []KV{}
		for k, v := range s.OptionsKV {
			kv = append(kv, KV{k, v})
		}
		type Alias SelectSelector
		o := struct {
			*Alias
			Options []KV `json:"options"`
		}{
			Alias:   (*Alias)(s),
			Options: kv,
		}
		return json.Marshal(o)
	}
	return json.Marshal(s)
}

type SelectSelectorMode string

const (
	SelectSelectorModeButtons  SelectSelectorMode = "list"
	SelectSelectorModeDropdown SelectSelectorMode = "dropdown"
)

type StateSelector struct {
	EntityID string `json:"entity_id"`
}

type TargetSelector struct {
	Device []SelectorDeviceFilter
	Entity []SelectorEntityFilter
}

type TextSelector struct {
	Multiline    bool             `json:"multiline"`
	Prefix       string           `json:"prefix,omitempty"`
	Suffix       string           `json:"suffix,omitempty"`
	Type         TextSelectorType `json:"type,omitempty"`
	Autocomplete string           `json:"autocomplete,omitempty"` // HTML autocomplete attribute
}

type TextSelectorType string

const (
	TextSelectorTypeColor         TextSelectorType = "color"
	TextSelectorTypeDate          TextSelectorType = "date"
	TextSelectorTypeDateTimeLocal TextSelectorType = "datetime-local"
	TextSelectorTypeEmail         TextSelectorType = "email"
	TextSelectorTypeMonth         TextSelectorType = "month"
	TextSelectorTypeNumber        TextSelectorType = "number"
	TextSelectorTypePassword      TextSelectorType = "password"
	TextSelectorTypeSearch        TextSelectorType = "search"
	TextSelectorTypeTel           TextSelectorType = "tel"
	TextSelectorTypeText          TextSelectorType = "text"
	TextSelectorTypeTime          TextSelectorType = "time"
	TextSelectorTypeURL           TextSelectorType = "url"
	TextSelectorTypeWeek          TextSelectorType = "week"
)

type ThemeSelector struct {
	IncludeDefault bool `json:"include_default"`
}

//

type SelectorDeviceFilter struct {
	Integration  string
	Manufacturer string
	Model        string
}

type SelectorEntityFilter struct {
	Integration       string
	Domain            string
	DeviceClass       []string
	SupportedFeatures []SupportedFeature
}

type ServiceFieldFilter struct {
	// one or the other, not both.
	Attribute         map[string]interface{} `json:"attribute,omitempty"`
	SupportedFeatures []SupportedFeature     `json:"supported_features,omitempty"`
}

//
