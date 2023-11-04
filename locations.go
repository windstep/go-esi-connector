package go_esi_connector

import (
	"encoding/json"
	"fmt"
	"time"
)

type Location struct {
	SolarSystemId int32 `json:"solar_system_id"`
	StationId     int32 `json:"station_id,omitempty"`
	StructureId   int32 `json:"structure_id,omitempty"`
}

type CharacterStatus struct {
	LastLogin  time.Time `json:"last_login,omitempty"`
	LastLogout time.Time `json:"last_logout,omitempty"`
	Logins     int32     `json:"logins,omitempty"`
	Online     bool      `json:"online"`
}

type CharacterShip struct {
	ShipItemId int64  `json:"ship_item_id"`
	ShipName   string `json:"ship_name"`
	ShipTypeId int32  `json:"ship_type_id"`
}

func (esi Client) GetCharacterLocation(characterId uint32, token string) (*Location, error) {
	body, _, err := esi.getAuth(fmt.Sprintf("/v2/characters/%d/location", characterId), token)
	if err != nil {
		return nil, err
	}

	var location Location
	if err = json.Unmarshal(body, &location); err != nil {
		return nil, err
	}

	return &location, nil
}

func (esi Client) GetCharacterOnline(characterId int32, token string) (*CharacterStatus, error) {
	body, _, err := esi.getAuth(fmt.Sprintf("/v3/characters/%d/online", characterId), token)
	if err != nil {
		return nil, err
	}

	var onlineStatus CharacterStatus
	if err = json.Unmarshal(body, &onlineStatus); err != nil {
		return nil, err
	}

	return &onlineStatus, nil
}

func (esi Client) GetCharacterShip(characterId int32, token string) (*CharacterShip, error) {
	body, _, err := esi.getAuth(fmt.Sprintf("/v2/characters/%d/ship", characterId), token)
	if err != nil {
		return nil, err
	}

	var ship CharacterShip
	if err = json.Unmarshal(body, &ship); err != nil {
		return nil, err
	}

	return &ship, nil
}
