package types

//ABILITIES
//GET /api/v2/ability/{id or name}/

//AbilityEffectChange holds the previous effect of this ability listed in different languages.
type AbilityEffectChange struct {
	EffectEntries []Effect         `json:"effect_entries"` //The previous effect of this ability listed in different languages.
	VersionGroups NamedAPIResource `json:"version_groups"` //The version group in which the previous effect of this ability originated.
}

//AbilityPokemon has the localized flavor for an API resource in a specific language.
type AbilityPokemon struct {
	IsHidden bool             `json:"is_hidden"` //Whether or not this a hidden ability for the referenced Pokémon.
	Slot     uint8            `json:"slot"`      //Pokémon have 3 ability 'slots' which hold references to possible abilities they could have. This is the slot of this ability for the referenced pokemon.
	Pokemon  NamedAPIResource `json:"pokemon"`   //The Pokémon this ability could belong to.
}

//AbilityFlavorText has The localized flavor text
type AbilityFlavorText struct {
	FlavorText   string           `json:"flavor_text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

//Ability holds ability info regarding name, id, gen, pokemon, and more.
type Ability struct {
	ID                uint16                `json:"id"`
	Name              string                `json:"name"`
	IsMainSeries      bool                  `json:"is_main_series"`
	Generation        NamedAPIResource      `json:"generation"`
	Names             []Name                `json:"names"`
	EffectEntries     []Effect              `json:"effect_entries"`
	EffectChange      []AbilityEffectChange `json:"effect_change"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}

//CHARACTERISTICS
//GET /api/v2/characteristic/{id}/

//Characteristic indicate which stat contains a Pokémon's highest IV. A Pokémon's Characteristic is determined by the remainder of its highest IV divided by 5 (gene_modulo). Check out Bulbapedia for greater detail.
type Characteristic struct {
	ID             uint16  `json:"id"`
	GeneModulo     uint8   `json:"gene_modulo"`
	PossibleValues []uint8 `json:"possible_values"`
}

//EGG GROUPS
//GET /api/v2/egg-group/{id or name}/

//EggGroup are categories which determine which Pokémon are able to interbreed. Pokémon may belong to either one or two Egg Groups. if - name - names - pokemonSpecies
type EggGroup struct {
	ID             uint16             `json:"id"`
	Name           string             `json:"name"`
	Names          []Name             `json:"names"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

//GENDERS
//GET /api/v2/gender/{id or name}/

//PokemonSpeciesGender defines the chance of a pokemon being a certain gender | rate - pokemonSpecies
type PokemonSpeciesGender struct {
	Rate           uint16           `json:"rate"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

//Gender holds details on pokemon gender | id - name - pokemonSpeciesDetails - requiredForEvolution
type Gender struct {
	ID                    uint16                 `json:"id"`
	Name                  string                 `json:"name"`
	PokemonSpeciesDetails []PokemonSpeciesGender `json:"pokemon_species_details"`
	RequiredForEvolution  []NamedAPIResource     `json:"required_for_evolution"`
}

//GROWTH RATES
//GET /api/v2/growth-rate/{id or name}/

//GrowthRateExperienceLevel indicates the level gained and exp required
type GrowthRateExperienceLevel struct {
	Level      uint8  `json:"level"`
	Experience uint32 `json:"experience"`
}

type GrowthRate struct {
	ID             uint16                      `json:"id"`
	Name           string                      `json:"name"`
	Formula        string                      `json:"formula"`
	Description    []Description               `json:"description"`
	Levels         []GrowthRateExperienceLevel `json:"levels"`
	PokemonSpecies []NamedAPIResource          `json:"pokemon_species"`
}

//NATURES
//GET /api/v2/nature/{id or name}/

type NatureStatChange struct {
	MaxChange      uint8            `json:"max_change"`
	PokeathlonStat NamedAPIResource `json:"pokeathlon_stat"`
}

type MoveBattleStylePreference struct {
	LowHpPreference  uint8            `json:"low_hp_preference"`
	HighHpPreference uint8            `json:"high_hp_preference"`
	MoveBattleStyle  NamedAPIResource `json:"move_battle_style"`
}

type Nature struct {
	ID                         uint16                      `json:"id"`
	Name                       string                      `json:"name"`
	DecreasedStat              NamedAPIResource            `json:"decreased_stat"`
	IncreasedStat              NamedAPIResource            `json:"increased_stat"`
	HatesFlavor                NamedAPIResource            `json:"hates_flavor"`
	LikesFlavor                NamedAPIResource            `json:"likes_flavor"`
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preference"`
	Names                      []Name                      `json:"names"`
}

//POKEATHLON STATS
//GET /api/v2/pokeathlon-stat/{id or name}/

type NaturePokeathlonStatAffect struct {
	MaxChange uint8            `json:"max_change"`
	Nature    NamedAPIResource `json:"nature"`
}

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect `json:"incrase"`
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}

type PokeathlonStat struct {
	ID               uint16                         `json:"id"`
	Name             string                         `json:"name"`
	Names            []Name                         `json:"names"`
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

//POKEMON
//GET /api/v2/pokemon/{id or name}/

type Pokemon struct {
	ID                     uint16             `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         uint16             `json:"base_experience"`
	Height                 uint16             `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  uint16             `json:"order"`
	Weight                 uint16             `json:"weight"`
	Abilities              []PokemonAbility   `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	Sprites                PokemonSprites     `json:"sprites"`
	Species                NamedAPIResource   `json:"species"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     uint8            `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	Slot uint8            `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonHeldItem struct {
	Item           NamedAPIResource         `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity  uint8            `json:"rarity"`
}

type PokemonMove struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  uint8            `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	Effort   uint8            `json:"effort"`
	BaseStat uint8            `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shinyFemale"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type LocationAreaEncounter struct {
	LocationArea   NamedAPIResource         `json:"location_area"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

//POKEMON COLORS
//GET /api/v2/pokemon-color/{id or name}/

type PokemonColor struct {
	ID             uint16             `json:"id"`
	Name           string             `json:"name"`
	Names          []Name             `json:"names"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

//POKEMON FORMS
//GET /api/v2/pokemon-form/{id or name}/

type PokemonForm struct {
	ID           uint16             `json:"is"`
	Name         string             `json:"name"`
	Order        uint16             `json:"order"`
	FormOrder    uint16             `json:"form_order"`
	IsDefault    bool               `json:"is_default"`
	IsBattleOnly bool               `json:"is_battle_only"`
	IsMega       bool               `json:"is_mega"`
	FormName     string             `json:"form_name"`
	Pokemon      NamedAPIResource   `json:"pokemon"`
	Sprites      PokemonFormSprites `json:"sprites"`
	VersionGroup NamedAPIResource   `json:"version_group"`
	Names        []Name             `json:"names"`
	FormNames    []Name             `json:"form_names"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}

//POKEMON HABITATS
//GET /api/v2/pokemon-habitat/{id or name}/

type PokemonHabitat struct {
	ID             uint16           `json:"id"`
	Name           string           `json:"name"`
	Names          []Name           `json:"names"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

//POKEMON SHAPES
//GET /api/v2/pokemon-shape/{id or name}/

type PokemonShape struct {
	ID             uint16           `json:"id"`
	Name           string           `json:"name"`
	AwesomeNames   []AwesomeName    `json:"awesome_names"`
	Names          []Name           `json:"name"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type AwesomeName struct {
	AwesomeName string           `json:"awesome_name"`
	Language    NamedAPIResource `json:"language"`
}

//POKEMON SPECIES
//GET /api/v2/pokemon-species/{id or name}/

type PokemonSpecies struct {
	ID                   uint16                   `json:"id"`
	Name                 string                   `json:"name"`
	Order                uint16                   `json:"order"`
	GenderRate           uint16                   `json:"gender_rate"`
	CaptureRate          uint8                    `json:"capture_rate"`
	BaseHappiness        uint8                    `json:"base_happiness"`
	IsBaby               bool                     `json:"is_baby"`
	HatchCounter         uint16                   `json:"hatch_counter"`
	HasGenderDifferences bool                     `json:"has_gender_differences"`
	FormsSwitchable      bool                     `json:"forms_switchable"`
	GrowthRate           NamedAPIResource         `json:"growth_rate"`
	PokedexNumbers       []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	EggGroups            []NamedAPIResource       `json:"egg_groups"`
	Color                NamedAPIResource         `json:"color"`
	Shape                NamedAPIResource         `json:"shape"`
	EvolvesFromSpecies   NamedAPIResource         `json:"evolves_from_species"`
	EvolutionChain       APIResource              `json:"evolution_chain"`
	Habitat              NamedAPIResource         `json:"habitat"`
	Generation           NamedAPIResource         `json:"generation"`
	Names                []Name                   `json:"names"`
	PalParkEncounters    []PalParkEncounterArea   `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText             `json:"flavor_text_entries"`
	FormDescription      []Description            `json:"form_description"`
	Genera               []Genus                  `json:"genera"`
	Varieties            []PokemonSpeciesVariety  `json:"varieties"`
}

type Genus struct {
	Genus    string           `json:"genus"`
	Language NamedAPIResource `json:"language"`
}

type PokemonSpeciesDexEntry struct {
	EntryNumber uint16           `json:"entry_number"`
	Pokedex     NamedAPIResource `json:"pokedex"`
}

type PalParkEncounterArea struct {
	BaseScore uint16           `json:"base_score"`
	Rate      uint8            `json:"rate"`
	Area      NamedAPIResource `json:"area"`
}

type PokemonSpeciesVariety struct {
	IsDefault bool             `json:"is_default"`
	Pokemon   NamedAPIResource `json:"pokemon"`
}

// STATS
//GET /api/v2/stat/{id or name}/

type Stat struct {
	ID               uint8                `json:"id"`
	Name             string               `json:"name"`
	GameIndex        uint16               `json:"game_index"`
	IsBattleOnly     bool                 `json:"is_battle_only"`
	AffectingMoves   MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	Characteristics  APIResource          `json:"characteristics"`
	MoveDamageClass  NamedAPIResource     `json:"move_damage_class"`
	Names            []Name               `json:"names"`
}

type MoveStatAffectSets struct {
	Increase []MoveStatAffect `json:"increase"`
	Decrease []MoveStatAffect `json:"decrease"`
}

type MoveStatAffect struct {
	Change uint8            `json:"change"`
	Move   NamedAPIResource `json:"move"`
}

type NatureStatAffectSets struct {
	Increase []NamedAPIResource `json:"increase"`
	Decrease []NamedAPIResource `json:"increase"`
}

//TYPES
//GET /api/v2/type/{id or name}/
type Type struct {
	ID              uint16                `json:"id"`
	Name            string                `json:"name"`
	DamageRelations TypeRelations         `json:"damage_relations"`
	GameIndices     []GenerationGameIndex `json:"game_indices"`
	Generation      NamedAPIResource      `json:"generation"`
	MoveDamageClass NamedAPIResource      `json:"move_damage_class"`
	Names           []Name                `json:"names"`
	Pokemon         []TypePokemon         `json:"pokemon"`
	Move            []NamedAPIResource    `json:"move"`
}

type TypePokemon struct {
	Slot    uint8            `json:"slot"`
	Pokemon NamedAPIResource `json:"pokemon"`
}

type TypeRelations struct {
	NoDamageTo       []NamedAPIResource `json:"no_damage_to"`
	HalfDamageTo     []NamedAPIResource `json:"half_damage_to"`
	DoubleDamageTo   []NamedAPIResource `json:"double_damage_to"`
	NoDamageFrom     []NamedAPIResource `json:"no_damage_from"`
	HalfDamageFrom   []NamedAPIResource `json:"half_damage_from"`
	DoubleDamageFrom []NamedAPIResource `json:"double_damage_from"`
}
