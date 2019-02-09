package types

type Item struct {
	ID                uint16                   `json:"id"`
	Name              string                   `json:"name"`
	Cost              uint                     `json:"cost"`
	FlingPower        uint8                    `json:"fling_power"`
	FlingEffect       NamedAPIResource         `json:"fling_effect"`
	Attributes        []NamedAPIResource       `json:"attributes"`
	Category          ItemCategory             `json:"category"`
	EffectEntries     []VerboseEffect          `json:"effect_entries"`
	FlavorTextEntries []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices       []GenerationGameIndex    `json:"game_indices"`
	Names             []Name                   `json:"names"`
	Sprites           ItemSprites              `json:"sprites"`
	HeldByPokemon     []ItemHolderPokemon      `json:"held_by_pokemon"`
	BabyTriggerFor    APIResource              `json:"baby_trigger_for"`
	Machines          []MachineVersionDetail   `json:"machines"`
}

type ItemSprites struct {
	Default string `json:"default"`
}

type ItemHolderPokemon struct {
	Pokemon        string                           `json:"pokemon"`
	VersionDetails []ItemHolderPokemonVersionDetail `json:"version_details"`
}

type ItemHolderPokemonVersionDetail struct {
	Rarity  string           `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type ItemAttribute struct {
	ID           uint16             `json:"id"`
	Name         string             `json:"name"`
	Items        []NamedAPIResource `json:"items"`
	Names        []Name             `json:"names"`
	Descriptions []Description      `json:"descriptions"`
}

type ItemCategory struct {
	ID     uint16             `json:"id"`
	Name   string             `json:"name"`
	Items  []NamedAPIResource `json:"items"`
	Names  []Name             `json:"names"`
	Pocket NamedAPIResource   `json:"pocket"`
}

type ItemFlingEffect struct {
	ID            uint16             `json:"id"`
	Name          string             `json:"name"`
	EffectEntries []Effect           `json:"effect_entries"`
	Items         []NamedAPIResource `json:"items"`
}

type ItemPocket struct {
	ID         uint16             `json:"id"`
	Name       string             `json:"name"`
	Categories []NamedAPIResource `json:"categories"`
	Names      []Name             `json:"names"`
}
