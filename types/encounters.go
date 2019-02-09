package types

type EncounterMethod struct {
	ID    uint16 `json:"id"`
	Name  string `json:"name"`
	Order uint16 `json:"order"`
	Names []Name `json:"names"`
}

type EncounterCondition struct {
	ID     uint16             `json:"id"`
	Name   string             `json:"name"`
	Names  []Name             `json:"names"`
	Values []NamedAPIResource `json:"values"`
}

type EncounterConditionValue struct {
	ID        uint16             `json:"id"`
	Name      string             `json:"name"`
	Condition []NamedAPIResource `json:"condition"`
	Names     []Name             `json:"names"`
}
