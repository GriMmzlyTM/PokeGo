package types

//BERRIES

type Berry struct {
	ID               uint16           `json:"id"`
	Name             string           `json:"name"`
	GrowthTime       uint8            `json:"growth_time"`
	MaxHarvest       uint8            `json:"max_harvest"`
	NaturalGiftPower uint8            `json:"natural_gift_power"`
	Size             uint16           `json:"size"`
	Smoothness       uint8            `json:"smoothness"`
	SoilDryness      uint16           `json:"soil_dryness"`
	Firmness         NamedAPIResource `json:"firmness"`
	Flavors          []BerryFlavorMap `json:"flavors"`
	Item             NamedAPIResource `json:"item"`
	NaturalGiftType  NamedAPIResource `json:"natural_gift_type"`
}

type BerryFlavorMap struct {
	Potency uint8            `json:"potency"`
	Flavor  NamedAPIResource `json:"flavor"`
}

type BerryFirmness struct {
	ID      uint16             `json:"id"`
	Name    string             `json:"name"`
	Berries []NamedAPIResource `json:"berries"`
	Names   []Name             `json:"names"`
}

type BerryFlavor struct {
	ID          uint16           `json:"id"`
	Name        string           `json:"name"`
	Berries     []FlavorBerryMap `json:"berries"`
	ContestType NamedAPIResource `json:"contest_type"`
	Names       []Name           `json:"names"`
}

type FlavorBerryMap struct {
	Potency uint8            `json:"potency"`
	Berry   NamedAPIResource `json:"berry"`
}
