package hass

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// CheckAPI checks if the API is running. It returns an error if it's not running.
func (a *Access) CheckAPI() error {
	response := struct {
		Message string `json:"message"`
	}{}
	err := a.httpGet(PathTypeAPI, "/", &response)
	if err != nil {
		return err
	}

	if response.Message == "" {
		return errors.New("hass: API is not running")
	}

	return nil
}

// State is the struct for an object state
type State struct {
	Attributes  StateAttributes `json:"attributes"`
	EntityID    string          `json:"entity_id"`
	LastChanged time.Time       `json:"last_changed"`
	LastUpdated time.Time       `json:"last_updated"`
	State       string          `json:"state"`
}

type StateAttributes map[string]interface{}

// States is an array of State objects
type States []State

// StateChange is used for changing state on an entity
type StateChange struct {
	EntityID string `json:"entityid"`
	State    string `json:"state"`
}

// GetDomain parses the Entity ID and returns the domain
func (s *State) GetDomain() string {
	return strings.TrimSuffix(strings.SplitAfter(s.EntityID, ".")[0], ".")
}

// FireEvent fires an event.
func (a *Access) FireEvent(eventType string, eventData interface{}) error {
	return a.httpPost(PathTypeAPI, "events/"+eventType, eventData, nil)
}

// CallService calls a service with a domain, service, and entity id.
func (a *Access) CallService(domain, service, entityID string) error {
	serviceData := struct {
		EntityID string `json:"entity_id"`
	}{entityID}
	return a.httpPost(PathTypeAPI, "services/"+domain+"/"+service, serviceData, nil)
}

// ListStates gets an array of state objects
func (a *Access) ListStates() (s States, err error) {
	var list States
	err = a.httpGet(PathTypeAPI, "states", &list)
	if err != nil {
		return States{}, err
	}
	return list, nil
}

// GetState retrieves one stateobject for the entity id
func (a *Access) GetState(id string) (s State, err error) {
	var state State
	err = a.httpGet(PathTypeAPI, "states/"+id, &state)
	if err != nil {
		return State{}, err
	}
	return state, nil
}

// FilterStates returns a list of states filtered by the list of domains
func (a *Access) FilterStates(domains ...string) (s States, err error) {
	list, err := a.ListStates()
	if err != nil {
		return States{}, err
	}
	for d := range list {
		for _, fdom := range domains {
			if fdom == list[d].GetDomain() {
				s = append(s, list[d])
			}
		}
		if err != nil {
			panic(err)
		}
	}

	return s, err
}

// ChangeState changes the state of a device
func (a *Access) ChangeState(id, state string) (s State, err error) {
	s.EntityID = id
	s.State = state
	err = a.httpPost(PathTypeAPI, "states/"+id, s, nil)
	return State{}, err
}

// ListAreas returns a list of areas
func (a *Access) ListAreas() (areas []string, err error) {
	var output string
	output, err = a.RenderTemplate(`{{ areas() | to_json }}`)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(output), &areas)
	if err != nil {
		return
	}

	return areas, nil
}

type AreaEntity struct {
	EntityID     string `json:"entity_id,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	State        string `json:"state,omitempty"`
	AreaID       string `json:"area_id,omitempty"`
	DeviceID     string `json:"device_id,omitempty"`
	DeviceClass  string `json:"device_class,omitempty"`
}

// ListAreaEntities returns a list of entities in an area
func (a *Access) ListAreaEntities(area string) (entities []AreaEntity, err error) {
	escapedArea := strings.Replace(area, `"`, `\"`, -1)
	var output string
	output, err = a.RenderTemplate(fmt.Sprintf(`
		{
			"entities": [
		{%%- for state in expand(area_entities("%s")) %%}
				{
					"entity_id": {{ state.entity_id | default("") | tojson }},
					"friendly_name": {{ state.attributes.friendly_name | default("") | tojson }},
					"state": {{ state.state | default("") | tojson }},
					"area_id": {{ area_id(state.entity_id) | default("") | tojson }},
					"device_id": {{ device_id(state.entity_id) | default("") | tojson }},
					"device_class": {{ state.attributes.device_class | default("") | tojson }}
				},
		{%%- endfor %%}
			]
		}
	`, escapedArea))
	if err != nil {
		return
	}

	output = regexp.MustCompile(`(\}),(\s*\]\s*\}\s*)$`).ReplaceAllString(output, `$1$2`)

	var response struct {
		Entities []AreaEntity `json:"entities"`
	}
	err = json.Unmarshal([]byte(output), &response)
	if err != nil {
		return
	}
	entities = response.Entities

	return entities, nil
}
