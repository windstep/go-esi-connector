package goesi

type LocationFlag string

const (
	AssetSafety                         LocationFlag = "AssetSafety"
	AutoFit                                          = "AutoFit"
	BoosterBay                                       = "BoosterBay"
	Cargo                                            = "Cargo"
	CorporationGoalDeliveries                        = "CorporationGoalDeliveries"
	CorpseBay                                        = "CorpseBay"
	Deliveries                                       = "Deliveries"
	DroneBay                                         = "DroneBay"
	FighterBay                                       = "FighterBay"
	FighterTube0                                     = "FighterTube0"
	FighterTube1                                     = "FighterTube1"
	FighterTube2                                     = "FighterTube2"
	FighterTube3                                     = "FighterTube3"
	FighterTube4                                     = "FighterTube4"
	FleetHangar                                      = "FleetHangar"
	FrigateEscapeBay                                 = "FrigateEscapeBay"
	Hangar                                           = "Hangar"
	HangarAll                                        = "HangarAll"
	HiSlot0                                          = "HiSlot0"
	HiSlot1                                          = "HiSlot1"
	HiSlot2                                          = "HiSlot2"
	HiSlot3                                          = "HiSlot3"
	HiSlot4                                          = "HiSlot4"
	HiSlot5                                          = "HiSlot5"
	HiSlot6                                          = "HiSlot6"
	HiSlot7                                          = "HiSlot7"
	HiddenModifiers                                  = "HiddenModifiers"
	Implant                                          = "Implant"
	LoSlot0                                          = "LoSlot0"
	LoSlot1                                          = "LoSlot1"
	LoSlot2                                          = "LoSlot2"
	LoSlot3                                          = "LoSlot3"
	LoSlot4                                          = "LoSlot4"
	LoSlot5                                          = "LoSlot5"
	LoSlot6                                          = "LoSlot6"
	LoSlot7                                          = "LoSlot7"
	Locked                                           = "Locked"
	MedSlot0                                         = "MedSlot0"
	MedSlot1                                         = "MedSlot1"
	MedSlot2                                         = "MedSlot2"
	MedSlot3                                         = "MedSlot3"
	MedSlot4                                         = "MedSlot4"
	MedSlot5                                         = "MedSlot5"
	MedSlot6                                         = "MedSlot6"
	MedSlot7                                         = "MedSlot7"
	MobileDepotHold                                  = "MobileDepotHold"
	QuafeBay                                         = "QuafeBay"
	RigSlot0                                         = "RigSlot0"
	RigSlot1                                         = "RigSlot1"
	RigSlot2                                         = "RigSlot2"
	RigSlot3                                         = "RigSlot3"
	RigSlot4                                         = "RigSlot4"
	RigSlot5                                         = "RigSlot5"
	RigSlot6                                         = "RigSlot6"
	RigSlot7                                         = "RigSlot7"
	ShipHangar                                       = "ShipHangar"
	Skill                                            = "Skill"
	SpecializedAmmoHold                              = "SpecializedAmmoHold"
	SpecializedAsteroidHold                          = "SpecializedAsteroidHold"
	SpecializedCommandCenterHold                     = "SpecializedCommandCenterHold"
	SpecializedFuelBay                               = "SpecializedFuelBay"
	SpecializedGasHold                               = "SpecializedGasHold"
	SpecializedIceHold                               = "SpecializedIceHold"
	SpecializedIndustrialShipHold                    = "SpecializedIndustrialShipHold"
	SpecializedLargeShipHold                         = "SpecializedLargeShipHold"
	SpecializedMaterialBay                           = "SpecializedMaterialBay"
	SpecializedMediumShipHold                        = "SpecializedMediumShipHold"
	SpecializedMineralHold                           = "SpecializedMineralHold"
	SpecializedOreHold                               = "SpecializedOreHold"
	SpecializedPlanetaryCommoditiesHold              = "SpecializedPlanetaryCommoditiesHold"
	SpecializedSalvageHold                           = "SpecializedSalvageHold"
	SpecializedShipHold                              = "SpecializedShipHold"
	SpecializedSmallShipHold                         = "SpecializedSmallShipHold"
	StructureDeedBay                                 = "StructureDeedBay"
	SubSystemBay                                     = "SubSystemBay"
	SubSystemSlot0                                   = "SubSystemSlot0"
	SubSystemSlot1                                   = "SubSystemSlot1"
	SubSystemSlot2                                   = "SubSystemSlot2"
	SubSystemSlot3                                   = "SubSystemSlot3"
	SubSystemSlot4                                   = "SubSystemSlot4"
	SubSystemSlot5                                   = "SubSystemSlot5"
	SubSystemSlot6                                   = "SubSystemSlot6"
	SubSystemSlot7                                   = "SubSystemSlot7"
	Unlocked                                         = "Unlocked"
	Wardrobe                                         = "Wardrobe"
)

type LocationType string

const (
	Station     LocationType = "station"
	SolarSystem              = "solar_system"
	Item                     = "item"
	Other                    = "other"
)

type Asset struct {
	IsBlueprintCopy bool         `json:"is_blueprint_copy"`
	IsSingleton     bool         `json:"is_singleton"`
	ItemId          int64        `json:"item_id"`
	LocationFlag    LocationFlag `json:"location_flag"`
	LocationId      int64        `json:"location_id"`
	LocationType    LocationType `json:"location_type"`
	Quantity        int32        `json:"quantity"`
	TypeId          int32        `json:"type_id"`
}
