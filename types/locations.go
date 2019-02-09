package types

type Location struct {
	ID          uint16                `json:"id"`
	Name        string                `json:"name"`
	Region      NamedAPIResource      `json:"region"`
	Names       []Name                `json:"names"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas       []NamedAPIResource    `json:"areas"`
}

type LocationArea struct {
	ID                   uint16                `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            uint16                `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    uint8            `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PalParkArea struct {
	ID                uint16                    `json:"id"`
	Name              string                    `json:"name"`
	Names             []Name                    `json:"names"`
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	BaseScore      uint8            `json:"base_score"`
	Rate           uint8            `json:"rate"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type Region struct {
	ID             uint16             `json:"id"`
	Locations      []NamedAPIResource `json:"locations"`
	Name           string             `json:"name"`
	Names          []Name             `json:"names"`
	MainGeneration NamedAPIResource   `json:"main_generation"`
	Pokedexes      []NamedAPIResource `json:"pokedexes"`
	VersionGroup   []NamedAPIResource `json:"version_group"`
}
