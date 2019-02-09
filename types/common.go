package types

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
	Generation NamedAPIResource `json:"generation"` //The generation relevent to this game index.
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
	Version   NamedAPIResource `json:"version"`    //The version relevent to this game index.
}

//VersionGroupFlavorText contains localized name and version group
type VersionGroupFlavorText struct {
	Text         string           `json:"text"`          //The localized name for an API resource in a specific language.
	Language     NamedAPIResource `json:"language"`      //The language this name is in.
	VersionGroup NamedAPIResource `json:"version_group"` //The version group which uses this flavor text.
}

type Language struct {
	ID       uint16 `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	ISO639   string `json:"iso639"`
	ISO3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}

type APIResourceList struct {
	Count    uint16        `json:"count"`
	Next     string        `json:"next"`
	Previous bool          `json:"previous"`
	Result   []APIResource `json:"results"`
}

type NamedAPIResourceList struct {
	Count    uint16             `json:"count"`
	Next     string             `json:"next"`
	Previous bool               `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}
