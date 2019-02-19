package pokego

import "fmt"

//COMMON

//API is the core parent struct for all data received from API calls.
type API struct {
	ID    uint16 `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}

//IsCached checks if the object name exists anywhere in the cache. Will return true if one or more
//files exist in the cache with this objects name.
func (c *API) IsCached() (bool, error) {
	loc := fmt.Sprintf("%s*/%s*", ops.CacheDir, c.Name)
	return checkCache(loc)
}

//APIResource holds "url" which leads to specified data
type APIResource struct {
	URL string `json:"url"` //The URL of the referenced resource.
}

//NamedAPIResource holds a name and url which leads to specified data. I.E. A pokemon, move, type, etc.
type NamedAPIResource struct {
	Name string `json:"name"` //The name of the referenced resource
	URL  string `json:"url"`  //The URL of the resource
}

//Description holds the description of an item, as well as the language the description is in.
type Description struct {
	Description string           `json:"description"` //The localized description for an API resource in a specific language.
	Language    NamedAPIResource `json:"language"`    //The language this name is in
}

//Effect holds the localized effect text for an API resource in a specific language
type Effect struct {
	Effect   string           `json:"effect"`   //The localized effect text for an API resource in a specific language.
	Language NamedAPIResource `json:"language"` //The language this effect is in.
}

//Encounter holds info regarding pokemon encounters.
type Encounter struct {
	MinLevel        uint8              `json:"min_level"`        //The lowest level the Pokémon could be encountered at.
	MaxLevel        uint8              `json:"max_level"`        //The highest level the Pokémon could be encountered at.
	ConditionValues []NamedAPIResource `json:"condition_values"` //A list of condition values that must be in effect for this encounter to occur.
	Chance          uint8              `json:"chance"`           //Percent chance that this encounter will occur.
	Method          NamedAPIResource   `json:"method"`           //The method by which this encounter happens.
}

//FlavorText holds the localized flavor text in a specific language.
type FlavorText struct {
	FlavorText string           `json:"flavor_text"` //The localized flavor text for an API resource in a specific language.
	Language   NamedAPIResource `json:"language"`    //The language this name is in.
	Version    NamedAPIResource `json:"version"`     //The game version this flavor text is extracted from.
}

//GenerationGameIndex holds the ID of the resource in game ID
type GenerationGameIndex struct {
	GameIndex  uint16           `json:"game_index"` //The internal id of an API resource within game data.
	Generation NamedAPIResource `json:"generation"` //The generation relevant to this game index.
}

//MachineVersionDetail holds the machine that teaches a move
type MachineVersionDetail struct {
	Machine      APIResource      `json:"machine"`       //The machine that teaches a move from an item.
	VersionGroup NamedAPIResource `json:"version_group"` //The version group of this specific machine.
}

//Name holds the localized name of the api resource
type Name struct {
	Name     string           `json:"name"`     //The localized name for an API resource in a specific language.
	Language NamedAPIResource `json:"language"` //The language this name is in.
}

//VerboseEffect holds a verbose version of the effect text
type VerboseEffect struct {
	Effect      string           `json:"effect"`       //The localized effect text for an API resource in a specific language.
	ShortEffect string           `json:"short_effect"` //The localized effect text in brief.
	Language    NamedAPIResource `json:"language"`     //The language this effect is in.
}

//VersionEncounterDetail contains version in which encounter occurs, and chance of encounter
type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`           //The game version this encounter happens in.
	MaxChance        uint8            `json:"max_chance"`        //The total percentage of all encounter potential.
	EncounterDetails []Encounter      `json:"encounter_details"` //A list of encounters and their specifics.
}

//VersionGameIndex contains internal ID of the resource in game data
type VersionGameIndex struct {
	GameIndex uint32           `json:"game_index"` //The internal id of an API resource within game data.
	Version   NamedAPIResource `json:"version"`    //The version relevant to this game index.
}

//VersionGroupFlavorText contains localized name and version group
type VersionGroupFlavorText struct {
	Text         string           `json:"text"`          //The localized name for an API resource in a specific language.
	Language     NamedAPIResource `json:"language"`      //The language this name is in.
	VersionGroup NamedAPIResource `json:"version_group"` //The version group which uses this flavor text.
}

//Language contains language data, generally used to define what language text is in.
type Language struct {
	API
	Official bool   `json:"official"`
	ISO639   string `json:"iso639"`
	ISO3166  string `json:"iso3166"`
}

//Get method allowing you to override data within struct. Provide it the name of the data you want to use.
//For example "pikachu", "hyper-beam", "hondew"
//This WILL NOT override type. Will only find data of its set type.
func (c *Language) Get(name string) error {
	err := callAPI("language/", name, c)
	return err
}

//APIResourceList contains node linked list data for anonymouse APIResources
type APIResourceList struct {
	Count    uint16        `json:"count"`
	Next     string        `json:"next"`
	Previous bool          `json:"previous"`
	Result   []APIResource `json:"results"`
}

//NamedAPIResourceList contains node linked list data for named API resources
type NamedAPIResourceList struct {
	Count    uint16             `json:"count"`
	Next     string             `json:"next"`
	Previous bool               `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

//BERRIES

//Berry contains data for berries within the game. Check Godocs or Pokeapi docs for structure info.
type Berry struct {
	API
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

//Get method allowing you to override data within struct. Provide it the name of the data you want to use.
//For example "pikachu", "hyper-beam", "hondew"
//This WILL NOT override type. Will only find data of its set type.
func (c *Berry) Get(name string) error {
	err := callAPI("berry/", name, c)
	return err
}

//BerryFlavorMap is used by the Berry struct, and contains Potency(uint8) and Flavor (NamedApiResource)-BerryFlavor
type BerryFlavorMap struct {
	Potency uint8            `json:"potency"`
	Flavor  NamedAPIResource `json:"flavor"`
}

//BerryFirmness contains an array of berries that match the name. This array holds NamedAPIResources to berries.
//You can use the Name of said NamedAPIResources to get that data. STRUCT TYPE: Berry
type BerryFirmness struct {
	API
	Berries []NamedAPIResource `json:"berries"`
}

//Get method allowing you to override data within struct. Provide it the name of the data you want to use.
//For example "pikachu", "hyper-beam", "hondew"
//This WILL NOT override type. Will only find data of its set type.
func (c *BerryFirmness) Get(name string) error {
	err := callAPI("berry-firmness/", name, c)
	return err
}

//BerryFlavor contains a FlavorBerryMap array, and a ContestType NamedApiResource (Call ContestType API using NamedApiResource name)
type BerryFlavor struct {
	API
	Berries     []FlavorBerryMap `json:"berries"`
	ContestType NamedAPIResource `json:"contest_type"`
}

//Get method allowing you to override data within struct. Provide it the name of the data you want to use.
//For example "pikachu", "hyper-beam", "hondew"
//This WILL NOT override type. Will only find data of its set type.
func (c *BerryFlavor) Get(name string) error {
	err := callAPI("berry-flavor/", name, c)
	return err
}

//FlavorBerryMap contains potency and berry (NamedApiResource)
type FlavorBerryMap struct {
	Potency uint8            `json:"potency"`
	Berry   NamedAPIResource `json:"berry"`
}

//CONTESTS

type ContestType struct {
	API
	BerryFlavor NamedAPIResource `json:"berry_flavor"`
}

func (c *ContestType) Get(name string) error {
	err := callAPI("contest-type/", name, c)
	return err
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

func (c *ContestEffect) Get(name string) error {
	err := callAPI("contest-effect/", name, c)
	return err
}

type SuperContestEffect struct {
	ID                uint16             `json:"id"`
	Appeal            uint8              `json:"appeal"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries"`
	Moves             []NamedAPIResource `json:"moves"`
}

func (c *SuperContestEffect) Get(ID string) error {
	err := callAPI("super-contest-effect/", ID, c)
	return err
}

//ENCOUNTERS

type EncounterMethod struct {
	API
	Order uint16 `json:"order"`
}

func (c *EncounterMethod) Get(name string) error {
	err := callAPI("encounter-method/", name, c)
	return err
}

type EncounterCondition struct {
	API
	Values []NamedAPIResource `json:"values"`
}

func (c *Encounter) Get(name string) error {
	err := callAPI("encounter-condition/", name, c)
	return err
}

type EncounterConditionValue struct {
	API
	Condition []NamedAPIResource `json:"condition"`
}

func (c *EncounterConditionValue) Get(name string) error {
	err := callAPI("encounter-condition-value/", name, c)
	return err
}

//EVOLUTION

type EvolutionChain struct {
	ID              uint16           `json:"id"`
	BabyTriggerItem NamedAPIResource `json:"baby_trigger_item"`
	Chain           ChainLink        `json:"chain"`
}

func (c *EvolutionChain) Get(ID string) error {
	err := callAPI("evolution-chain/", ID, c)
	return err
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
	API
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

func (c *EvolutionTrigger) Get(name string) error {
	err := callAPI("evolution-trigger/", name, c)
	return err
}

//GAMES

type Generation struct {
	API
	Abilities      []NamedAPIResource `json:"abilities"`
	MainRegion     NamedAPIResource   `json:"main_region"`
	Moves          []NamedAPIResource `json:"moves"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
	Types          []NamedAPIResource `json:"types"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

func (c *Generation) Get(name string) error {
	err := callAPI("generation/", name, c)
	return err
}

//POKEDEX

type Pokedex struct {
	API
	IsMainSeries   bool               `json:"is_main_series"`
	Descriptions   []Description      `json:"descriptions"`
	PokemonEntries []PokemonEntry     `json:"pokemon_entries"`
	Region         NamedAPIResource   `json:"region"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

func (c *Pokedex) Get(name string) error {
	err := callAPI("pokedex/", name, c)
	return err
}

type PokemonEntry struct {
	EntryNumber    uint16           `json:"entry_number"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

//VERSION

type Version struct {
	API
	VersionGroup NamedAPIResource `json:"version_group"`
}

func (c *Version) Get(name string) error {
	err := callAPI("version/", name, c)
	return err
}

type VersionGroup struct {
	API
	Order            uint16             `json:"order"`
	Generation       NamedAPIResource   `json:"generation"`
	MoveLearnMethods []NamedAPIResource `json:"move_learn_methods"`
	Pokedexes        []NamedAPIResource `json:"pokedexes"`
	Regions          []NamedAPIResource `json:"regions"`
	Versions         []NamedAPIResource `json:"versions"`
}

func (c *VersionGroup) Get(name string) error {
	err := callAPI("version-group/", name, c)
	return err
}

//ITEMS

type Item struct {
	API
	Cost              uint                     `json:"cost"`
	FlingPower        uint8                    `json:"fling_power"`
	FlingEffect       NamedAPIResource         `json:"fling_effect"`
	Attributes        []NamedAPIResource       `json:"attributes"`
	Category          ItemCategory             `json:"category"`
	EffectEntries     []VerboseEffect          `json:"effect_entries"`
	FlavorTextEntries []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices       []GenerationGameIndex    `json:"game_indices"`
	Sprites           ItemSprites              `json:"sprites"`
	HeldByPokemon     []ItemHolderPokemon      `json:"held_by_pokemon"`
	BabyTriggerFor    APIResource              `json:"baby_trigger_for"`
	Machines          []MachineVersionDetail   `json:"machines"`
}

func (c *Item) Get(name string) error {
	err := callAPI("item/", name, c)
	return err
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
	API
	Items        []NamedAPIResource `json:"items"`
	Descriptions []Description      `json:"descriptions"`
}

func (c *ItemAttribute) Get(name string) error {
	err := callAPI("item-attribute/", name, c)
	return err
}

type ItemCategory struct {
	API
	Items  []NamedAPIResource `json:"items"`
	Pocket NamedAPIResource   `json:"pocket"`
}

func (c *ItemCategory) Get(name string) error {
	err := callAPI("item-category/", name, c)
	return err
}

type ItemFlingEffect struct {
	API
	EffectEntries []Effect           `json:"effect_entries"`
	Items         []NamedAPIResource `json:"items"`
}

func (c *ItemFlingEffect) Get(name string) error {
	err := callAPI("item-fling-effect/", name, c)
	return err
}

type ItemPocket struct {
	API
	Categories []NamedAPIResource `json:"categories"`
}

func (c *ItemPocket) Get(name string) error {
	err := callAPI("item-pocket/", name, c)
	return err
}

//LOCATION

type Location struct {
	API
	Region      NamedAPIResource      `json:"region"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas       []NamedAPIResource    `json:"areas"`
}

func (c *Location) Get(name string) error {
	err := callAPI("location/", name, c)
	return err
}

type LocationArea struct {
	API
	GameIndex            uint16                `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

func (c *LocationArea) Get(name string) error {
	err := callAPI("location-area/", name, c)
	return err
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
	API
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

func (c *PalParkArea) Get(name string) error {
	err := callAPI("pal-park-area/", name, c)
	return err
}

type PalParkEncounterSpecies struct {
	BaseScore      uint8            `json:"base_score"`
	Rate           uint8            `json:"rate"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type Region struct {
	API
	Locations      []NamedAPIResource `json:"locations"`
	MainGeneration NamedAPIResource   `json:"main_generation"`
	Pokedexes      []NamedAPIResource `json:"pokedexes"`
	VersionGroup   []NamedAPIResource `json:"version_group"`
}

func (c *Region) Get(name string) error {
	err := callAPI("region/", name, c)
	return err
}

//MACHINE

type Machine struct {
	API
	Item         NamedAPIResource `json:"item"`
	Move         NamedAPIResource `json:"move"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

func (c *Machine) Get(ID string) error {
	err := callAPI("machine/", ID, c)
	return err
}

//MOVES

type Move struct {
	API
	Accuracy           uint8                  `json:"accuracy"`
	EffectChance       uint8                  `json:"effect_chance"`
	PP                 uint8                  `json:"pp"`
	Priority           int8                   `json:"priority"`
	Power              uint8                  `json:"power"`
	ContestCombos      ContestComboSets       `json:"contest_combos"`
	ContestType        NamedAPIResource       `json:"contest_type"`
	ContestEffect      APIResource            `json:"contest_effect"`
	DamageClass        NamedAPIResource       `json:"damage_class"`
	EffectEntries      []VerboseEffect        `json:"effect_entries"`
	EffectChanges      []AbilityEffectChange  `json:"effect_changes"`
	FlavorTextEntries  []MoveFlavorText       `json:"flavor_text_entries"`
	Generation         NamedAPIResource       `json:"generation"`
	Machines           []MachineVersionDetail `json:"machines"`
	Meta               MoveMetaData           `json:"meta"`
	PastValues         []PastMoveStatValues   `json:"past_values"`
	StatChanges        []MoveStatChange       `json:"stat_changes"`
	SuperContestEffect APIResource            `json:"super_contest_effect"`
	Target             NamedAPIResource       `json:"target"`
	Type               NamedAPIResource       `json:"type"`
}

func (c *Move) Get(name string) error {
	err := callAPI("move/", name, c)
	return err
}

type ContestComboSets struct {
	Normal ContestComboDetail `json:"normal"`
	Super  ContestComboDetail `json:"super"`
}

type ContestComboDetail struct {
	UseBefore []NamedAPIResource `json:"use_before"`
	UseAfter  []NamedAPIResource `json:"use_after"`
}

type MoveFlavorText struct {
	FlavorText   string           `json:"flavor_text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type MoveMetaData struct {
	Ailment       NamedAPIResource `json:"ailment"`
	Category      NamedAPIResource `json:"category"`
	MinHits       uint8            `json:"min_hits"`
	MaxHits       uint8            `json:"max_hits"`
	MinTurns      uint8            `json:"min_turns"`
	MaxTurns      uint8            `json:"max_turns"`
	Drain         int8             `json:"drain"`
	Healing       uint8            `json:"healing"`
	CritRate      uint8            `json:"crit_rate"`
	AilmentChance uint8            `json:"ailment_chance"`
	FlinchChance  uint8            `json:"clinch_chance"`
	StatChance    uint8            `json:"stat_chance"`
}

type MoveStatChange struct {
	Change uint8            `json:"change"`
	Stat   NamedAPIResource `json:"stat"`
}

type PastMoveStatValues struct {
	Accuracy      uint8            `json:"accuracy"`
	EffectChance  uint8            `json:"effect_chance"`
	Power         uint8            `json:"power"`
	PP            uint8            `json:"pp"`
	EffectEntries []VerboseEffect  `json:"effect_entries"`
	Type          NamedAPIResource `json:"type"`
	VersionGroup  NamedAPIResource `json:"version_group"`
}

type MoveAilment struct {
	API
	Moves []NamedAPIResource `json:"moves"`
}

func (c *MoveAilment) Get(name string) error {
	err := callAPI("move-ailment/", name, c)
	return err
}

type MoveBattleStyle struct {
	API
}

func (c *MoveBattleStyle) Get(name string) error {
	err := callAPI("move-battle-style/", name, c)
	return err
}

type MoveCategory struct {
	API
	Moves        []NamedAPIResource `json:"moves"`
	Descriptions []Description      `json:"descriptions"`
}

func (c *MoveCategory) Get(name string) error {
	err := callAPI("move-category/", name, c)
	return err
}

type MoveDamageClass struct {
	API
	Descriptions []Description      `json:"descriptions"`
	Moves        []NamedAPIResource `json:"moves"`
}

func (c *MoveDamageClass) Get(name string) error {
	err := callAPI("move-damage-class/", name, c)
	return err
}

type MoveLearnMethod struct {
	API
	Descriptions  []Description      `json:"descriptions"`
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

func (c *MoveLearnMethod) Get(name string) error {
	err := callAPI("move-learn-method/", name, c)
	return err
}

type MoveTarget struct {
	API
	Descriptions []Description      `json:"descriptions"`
	Moves        []NamedAPIResource `json:"moves"`
}

func (c *MoveTarget) Get(name string) error {
	err := callAPI("move-target/", name, c)
	return err
}

//POKEMON

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
	API
	IsMainSeries      bool                  `json:"is_main_series"`
	Generation        NamedAPIResource      `json:"generation"`
	EffectEntries     []Effect              `json:"effect_entries"`
	EffectChange      []AbilityEffectChange `json:"effect_change"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}

func (c *Ability) Get(name string) error {
	err := callAPI("ability/", name, c)
	return err
}

//CHARACTERISTICS
//GET /api/v2/characteristic/{id}/

//Characteristic indicate which stat contains a Pokémon's highest IV. A Pokémon's Characteristic is determined by the remainder of its highest IV divided by 5 (gene_modulo). Check out Bulbapedia for greater detail.
type Characteristic struct {
	API
	GeneModulo     uint8   `json:"gene_modulo"`
	PossibleValues []uint8 `json:"possible_values"`
}

func (c *Characteristic) Get(name string) error {
	err := callAPI("characteristic/", name, c)
	return err
}

//EGG GROUPS
//GET /api/v2/egg-group/{id or name}/

//EggGroup are categories which determine which Pokémon are able to interbreed. Pokémon may belong to either one or two Egg Groups. if - name - names - pokemonSpecies
type EggGroup struct {
	API
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

func (c *EggGroup) Get(name string) error {
	err := callAPI("egg-group/", name, c)
	return err
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
	API
	PokemonSpeciesDetails []PokemonSpeciesGender `json:"pokemon_species_details"`
	RequiredForEvolution  []NamedAPIResource     `json:"required_for_evolution"`
}

func (c *Gender) Get(name string) error {
	err := callAPI("gender/", name, c)
	return err
}

//GROWTH RATES
//GET /api/v2/growth-rate/{id or name}/

//GrowthRateExperienceLevel indicates the level gained and exp required
type GrowthRateExperienceLevel struct {
	Level      uint8  `json:"level"`
	Experience uint32 `json:"experience"`
}

type GrowthRate struct {
	API
	Formula        string                      `json:"formula"`
	Description    []Description               `json:"description"`
	Levels         []GrowthRateExperienceLevel `json:"levels"`
	PokemonSpecies []NamedAPIResource          `json:"pokemon_species"`
}

func (c *GrowthRate) Get(name string) error {
	err := callAPI("growth-rate/", name, c)
	return err
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
	API
	DecreasedStat              NamedAPIResource            `json:"decreased_stat"`
	IncreasedStat              NamedAPIResource            `json:"increased_stat"`
	HatesFlavor                NamedAPIResource            `json:"hates_flavor"`
	LikesFlavor                NamedAPIResource            `json:"likes_flavor"`
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preference"`
}

func (c *Nature) Get(name string) error {
	err := callAPI("nature/", name, c)
	return err
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
	API
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

func (c *PokeathlonStat) Get(name string) error {
	err := callAPI("pokeathlon-stat/", name, c)
	return err
}

//POKEMON
//GET /api/v2/pokemon/{id or name}/

type Pokemon struct {
	API
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

func (c *Pokemon) Get(name string) error {
	err := callAPI("pokemon/", name, c)
	return err
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
	API
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

func (c *PokemonColor) Get(name string) error {
	err := callAPI("pokemon-color/", name, c)
	return err
}

//POKEMON FORMS
//GET /api/v2/pokemon-form/{id or name}/

type PokemonForm struct {
	API
	Order        uint16             `json:"order"`
	FormOrder    uint16             `json:"form_order"`
	IsDefault    bool               `json:"is_default"`
	IsBattleOnly bool               `json:"is_battle_only"`
	IsMega       bool               `json:"is_mega"`
	FormName     string             `json:"form_name"`
	Pokemon      NamedAPIResource   `json:"pokemon"`
	Sprites      PokemonFormSprites `json:"sprites"`
	VersionGroup NamedAPIResource   `json:"version_group"`
	FormNames    []Name             `json:"form_names"`
}

func (c *PokemonForm) Get(name string) error {
	err := callAPI("pokemon-form/", name, c)
	return err
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
	API
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

func (c *PokemonHabitat) Get(name string) error {
	err := callAPI("pokemon-habitat/", name, c)
	return err
}

//POKEMON SHAPES
//GET /api/v2/pokemon-shape/{id or name}/

type PokemonShape struct {
	API
	AwesomeNames   []AwesomeName    `json:"awesome_names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

func (c *PokemonShape) Get(name string) error {
	err := callAPI("pokemon-shape/", name, c)
	return err
}

type AwesomeName struct {
	AwesomeName string           `json:"awesome_name"`
	Language    NamedAPIResource `json:"language"`
}

//POKEMON SPECIES
//GET /api/v2/pokemon-species/{id or name}/

type PokemonSpecies struct {
	API
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
	PalParkEncounters    []PalParkEncounterArea   `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText             `json:"flavor_text_entries"`
	FormDescription      []Description            `json:"form_description"`
	Genera               []Genus                  `json:"genera"`
	Varieties            []PokemonSpeciesVariety  `json:"varieties"`
}

func (c *PokemonSpecies) Get(name string) error {
	err := callAPI("pokemon-species/", name, c)
	return err
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
	API
	GameIndex        uint16               `json:"game_index"`
	IsBattleOnly     bool                 `json:"is_battle_only"`
	AffectingMoves   MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	Characteristics  APIResource          `json:"characteristics"`
	MoveDamageClass  NamedAPIResource     `json:"move_damage_class"`
}

func (c *Stat) Get(name string) error {
	err := callAPI("stat/", name, c)
	return err
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
	Decrease []NamedAPIResource `json:"decrease"`
}

//GET /api/v2/type/{id or name}/
//Type
type Type struct {
	API
	DamageRelations TypeRelations         `json:"damage_relations"`
	GameIndices     []GenerationGameIndex `json:"game_indices"`
	Generation      NamedAPIResource      `json:"generation"`
	MoveDamageClass NamedAPIResource      `json:"move_damage_class"`
	Pokemon         []TypePokemon         `json:"pokemon"`
	Move            []NamedAPIResource    `json:"move"`
}

func (c *Type) Get(name string) error {
	err := callAPI("type/", name, c)
	return err
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
