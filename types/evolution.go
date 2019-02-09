package types

type EvolutionChain struct {
	ID              uint16           `json:"id"`
	BabyTriggerItem NamedAPIResource `json:"baby_trigger_item"`
	Chain           ChainLink        `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool              `json:"is_baby"`
	Species          NamedAPIResource  `json:"species"`
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo        []ChainLink       `json:"evolves_to"`
}

type EvolutionDetail struct {
	Item                  NamedAPIResource `json:"item"`
	Trigger               NamedAPIResource `json:"trigger"`
	Gender                uint8            `json:"gender"`
	HeldItem              NamedAPIResource `json:"held_item"`
	KnownMove             NamedAPIResource `json:"known_move"`
	KnownMoveType         NamedAPIResource `json:"known_move_type"`
	Location              NamedAPIResource `json:"location"`
	MinLevel              uint8            `json:"min_level"`
	MinHappiness          uint8            `json:"min_happiness"`
	MinBeauty             uint8            `json:"min_beauty"`
	MinAffection          uint8            `json:"min_affection"`
	NeedsOverworldRain    bool             `json:"needs_overworld_rain"`
	PartyType             NamedAPIResource `json:"party_type"`
	RelativePhysicalStats uint8            `json:"rekative_physical_stats"`
	TimeOfDay             string           `json:"time_of_day"`
	TradeSpecies          NamedAPIResource `json:"trade_species"`
	TurnUpsideDown        bool             `json:"turn_upside_down"`
}

type EvolutionTrigger struct {
	ID             uint16             `json:"id"`
	Name           string             `json:"name"`
	Names          []Name             `json:"names"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}
