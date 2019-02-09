package types

//CONTESTS

type ContestType struct {
	ID          uint16           `json:"id"`
	Name        string           `json:"name"`
	BerryFlavor NamedAPIResource `json:"berry_flavor"`
	Names       []ContestName    `json:"names"`
}

type ContestName struct {
	Name     string           `json:"name"`
	Color    string           `json:"color"`
	Language NamedAPIResource `json:"language"`
}

type ContestEffect struct {
	ID                uint16       `json:"id"`
	Appeal            uint8        `json:"appeal"`
	Jam               uint8        `json:"jam"`
	EffectEntries     []Effect     `json:"effects_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffect struct {
	ID                uint16             `json:"id"`
	Appeal            uint8              `json:"appeal"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries"`
	Moves             []NamedAPIResource `json:"moves"`
}
