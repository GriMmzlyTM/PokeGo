package types

type Generation struct {
	ID             uint16             `json:"id"`
	Name           string             `json:"name"`
	Abilities      []NamedAPIResource `json:"abilities"`
	Names          []Name             `json:"names"`
	MainRegion     NamedAPIResource   `json:"main_region"`
	Moves          []NamedAPIResource `json:"moves"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
	Types          []NamedAPIResource `json:"types"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

//POKEDEX

type Pokedex struct {
	ID             uint16             `json:"id"`
	Name           string             `json:"name"`
	IsMainSeries   bool               `json:"is_main_series"`
	Descriptions   []Description      `json:"descriptions"`
	Names          []Name             `json:"names"`
	PokemonEntries []PokemonEntry     `json:"pokemon_entries"`
	Region         NamedAPIResource   `json:"region"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

type PokemonEntry struct {
	EntryNumber    uint16           `json:"entry_number"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

//VERSION

type Version struct {
	ID           uint16           `json:"id"`
	Name         string           `json:"name"`
	Names        []Name           `json:"names"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type VersionGroup struct {
	ID               uint16             `json:"id"`
	Name             string             `json:"name"`
	Order            uint16             `json:"order"`
	Generation       NamedAPIResource   `json:"generation"`
	MoveLearnMethods []NamedAPIResource `json:"move_learn_methods"`
	Pokedexes        []NamedAPIResource `json:"pokedexes"`
	Regions          []NamedAPIResource `json:"regions"`
	Versions         []NamedAPIResource `json:"versions"`
}
