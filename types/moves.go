package types

//MOVES

type Move struct {
	ID                 uint16                 `json:"id"`
	Name               string                 `json:"name"`
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
	Names              []Name                 `json:"names"`
	PastValues         []PastMoveStatValues   `json:"past_values"`
	StatChanges        []MoveStatChange       `json:"stat_changes"`
	SuperContestEffect APIResource            `json:"super_contest_effect"`
	Target             NamedAPIResource       `json:"target"`
	Type               NamedAPIResource       `json:"type"`
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
	ID    uint16             `json:"is"`
	Name  string             `json:"name"`
	Moves []NamedAPIResource `json:"moves"`
	Names []Name             `json:"names"`
}

type MoveBattleStyle struct {
	ID    uint16 `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}

type MoveCategory struct {
	ID           uint16             `json:"id"`
	Name         string             `json:"name"`
	Moves        []NamedAPIResource `json:"moves"`
	Descriptions []Description      `json:"descriptions"`
}

type MoveDamageClass struct {
	ID           uint16             `json:"id"`
	Name         string             `json:"name"`
	Descriptions []Description      `json:"descriptions"`
	Moves        []NamedAPIResource `json:"moves"`
	Names        []Name             `json:"names"`
}

type MoveLearnMethod struct {
	ID            uint16             `json:"id"`
	Name          string             `json:"name"`
	Descriptions  []Description      `json:"descriptions"`
	Names         []Name             `json:"names"`
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

type MoveTarget struct {
	ID           uint16             `json:"id"`
	Name         string             `json:"name"`
	Descriptions []Description      `json:"descriptions"`
	Moves        []NamedAPIResource `json:"moves"`
	Names        []Name             `json:"names"`
}
