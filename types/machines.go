package types

type Machine struct {
	ID           uint16           `json:"id"`
	Item         NamedAPIResource `json:"item"`
	Move         NamedAPIResource `json:"move"`
	VersionGroup NamedAPIResource `json:"version_group"`
}
