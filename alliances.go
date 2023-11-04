package go_esi_connector

import (
	"encoding/json"
	"fmt"
	"time"
)

type Alliance struct {
	ID                    uint32
	CreatorCorporationId  uint32    `json:"creator_corporation_id"`  // ID of the corporation that created the alliance
	CreatorId             uint32    `json:"creator_id"`              // ID of the character that created the alliance
	DateFounded           time.Time `json:"date_founded"`            // date_founded string
	ExecutorCorporationId uint32    `json:"executor_corporation_id"` // the executor corporation ID, if this alliance is not closed
	FactionId             uint32    `json:"faction_id"`              // Faction ID this alliance is fighting for, if this alliance is enlisted in factional warfare
	Name                  string    `json:"name"`                    // the full name of the alliance
	Ticker                string    `json:"ticker"`                  // the short name of the alliance
}

type AllianceIcons struct {
	Px128x128 string `json:"px128x128"`
	Px64x64   string `json:"px64X64"`
}

func (esi Client) GetAlliancesIds() ([]uint32, error) {
	body, _, err := esi.get("/v2/alliances")
	if err != nil {
		return nil, err
	}

	var allianceIds []uint32
	if err = json.Unmarshal(body, &allianceIds); err != nil {
		return nil, err
	}

	return allianceIds, nil
}

func (esi Client) GetAlliance(allianceId uint32) (*Alliance, error) {
	body, _, err := esi.get(fmt.Sprintf("/v4/alliances/%d", allianceId))
	if err != nil {
		return nil, err
	}

	alliance := Alliance{ID: allianceId}
	if err = json.Unmarshal(body, &alliance); err != nil {
		return nil, err
	}

	return &alliance, nil
}

func (esi Client) GetAllianceCorporationIds(allianceId uint32) ([]uint32, error) {
	body, _, err := esi.get(fmt.Sprintf("/v2/alliances/%d/corporations", allianceId))
	if err != nil {
		return nil, err
	}

	var corporationsIds []uint32
	if err = json.Unmarshal(body, &corporationsIds); err != nil {
		return nil, err
	}

	return corporationsIds, nil
}

func (esi Client) GetAllianceIcons(allianceId uint32) (*AllianceIcons, error) {
	body, _, err := esi.get(fmt.Sprintf("/v4/alliances/%d/icons", allianceId))
	if err != nil {
		return nil, err
	}

	var allianceIcons AllianceIcons
	if err = json.Unmarshal(body, &allianceIcons); err != nil {
		return nil, err
	}

	return &allianceIcons, nil
}
