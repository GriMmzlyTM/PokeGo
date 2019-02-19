package pokego

//Berry returns a pointer a Berry type var
func GetBerry(name string) (*Berry, error) {
	berry := &Berry{}
	err := callAPI("berry/", name, berry)
	return berry, err
}

//BerryFirmness returns a pointer to a BerryFirmness var
func GetBerryFirmness(name string) (*BerryFirmness, error) {

	obj := &BerryFirmness{}
	err := callAPI("berry-firmness/", name, obj)
	return obj, err
}

//BerryFlavor returns a pointer.
func GetBerryFlavor(name string) (*BerryFlavor, error) {

	obj := &BerryFlavor{}
	err := callAPI("berry-flavor/", name, obj)
	return obj, err
}

func GetContestType(name string) (*ContestType, error) {

	obj := &ContestType{}
	err := callAPI("contest-type/", name, obj)
	return obj, err
}

func GetContestEffect(id string) (*ContestEffect, error) {

	obj := &ContestEffect{}
	err := callAPI("contest-effect/", id, obj)
	return obj, err
}

func GetSuperContestEffect(id string) (*SuperContestEffect, error) {

	obj := &SuperContestEffect{}
	err := callAPI("super-contest-effect/", id, obj)
	return obj, err
}

func GetEncounterMethod(name string) (*EncounterMethod, error) {

	obj := &EncounterMethod{}
	err := callAPI("encounter-method/", name, obj)
	return obj, err
}

func GetEncounterCondition(name string) (*EncounterCondition, error) {

	obj := &EncounterCondition{}
	err := callAPI("encounter-condition/", name, obj)
	return obj, err
}

func GetEncounterConditionValue(name string) (*EncounterConditionValue, error) {

	obj := &EncounterConditionValue{}
	err := callAPI("encounter-condition-value/", name, obj)
	return obj, err
}

func GetEvolutionChain(name string) (*EvolutionChain, error) {

	obj := &EvolutionChain{}
	err := callAPI("evolution-chain/", name, obj)
	return obj, err
}

func GetEvolutionTrigger(name string) (*EvolutionTrigger, error) {

	obj := &EvolutionTrigger{}
	err := callAPI("evolution-trigger/", name, obj)
	return obj, err
}

func GetGeneration(name string) (*Generation, error) {

	obj := &Generation{}
	err := callAPI("generation/", name, obj)
	return obj, err
}

func GetPokedex(name string) (*Pokedex, error) {

	obj := &Pokedex{}
	err := callAPI("pokedex/", name, obj)
	return obj, err
}

func GetVersion(name string) (*Version, error) {

	obj := &Version{}
	err := callAPI("version/", name, obj)
	return obj, err
}

func GetVersionGroup(name string) (*VersionGroup, error) {

	obj := &VersionGroup{}
	err := callAPI("version-group/", name, obj)
	return obj, err
}

func GetItem(name string) (*Item, error) {

	obj := &Item{}
	err := callAPI("item/", name, obj)
	return obj, err
}

func GetItemAttribute(name string) (*ItemAttribute, error) {

	obj := &ItemAttribute{}
	err := callAPI("item-attribute/", name, obj)
	return obj, err
}

func GetItemCategory(name string) (*ItemCategory, error) {

	obj := &ItemCategory{}
	err := callAPI("item-category/", name, obj)
	return obj, err
}

func GetItemFlingEffect(name string) (*ItemFlingEffect, error) {

	obj := &ItemFlingEffect{}
	err := callAPI("item-fling-effect/", name, obj)
	return obj, err
}

func GetItemPocket(name string) (*ItemPocket, error) {

	obj := &ItemPocket{}
	err := callAPI("item-pocket/", name, obj)
	return obj, err
}

func GetLocation(name string) (*Location, error) {

	obj := &Location{}
	err := callAPI("location/", name, obj)
	return obj, err
}

func GetLocationArea(name string) (*LocationArea, error) {

	obj := &LocationArea{}
	err := callAPI("location-area/", name, obj)
	return obj, err
}

func GetPalParkArea(name string) (*PalParkArea, error) {

	obj := &PalParkArea{}
	err := callAPI("pal-park-area/", name, obj)
	return obj, err
}

func GetRegion(name string) (*Region, error) {

	obj := &Region{}
	err := callAPI("region/", name, obj)
	return obj, err
}

func GetMachine(name string) (*Machine, error) {

	obj := &Machine{}
	err := callAPI("machine/", name, obj)
	return obj, err
}

func GetMove(name string) (*Move, error) {

	obj := &Move{}
	err := callAPI("move/", name, obj)
	return obj, err
}

func GetMoveAilment(name string) (*MoveAilment, error) {

	obj := &MoveAilment{}
	err := callAPI("move-ailment/", name, obj)
	return obj, err
}

func GetMoveBattleStyle(name string) (*MoveBattleStyle, error) {

	obj := &MoveBattleStyle{}
	err := callAPI("move-battle-style/", name, obj)
	return obj, err
}

func GetMoveCategory(name string) (*MoveCategory, error) {

	obj := &MoveCategory{}
	err := callAPI("move-category/", name, obj)
	return obj, err
}

func GetMoveDamageClass(name string) (*MoveDamageClass, error) {

	obj := &MoveDamageClass{}
	err := callAPI("move-damage-class/", name, obj)
	return obj, err
}

func GetMoveLearnMethod(name string) (*MoveLearnMethod, error) {

	obj := &MoveLearnMethod{}
	err := callAPI("move-learn-method/", name, obj)
	return obj, err
}

func GetMoveTarget(name string) (*MoveTarget, error) {

	obj := &MoveTarget{}
	err := callAPI("move-target/", name, obj)
	return obj, err
}

func GetAbility(name string) (*Ability, error) {

	obj := &Ability{}
	err := callAPI("ability/", name, obj)
	return obj, err
}

func GetCharacteristic(name string) (*Characteristic, error) {

	obj := &Characteristic{}
	err := callAPI("characteristic/", name, obj)
	return obj, err
}

func GetEggGroup(name string) (*EggGroup, error) {

	obj := &EggGroup{}
	err := callAPI("egg-group/", name, obj)
	return obj, err
}

func GetGender(name string) (*Gender, error) {

	obj := &Gender{}
	err := callAPI("gender/", name, obj)
	return obj, err
}

func GetGrowthRate(name string) (*GrowthRate, error) {

	obj := &GrowthRate{}
	err := callAPI("growth-rate/", name, obj)
	return obj, err
}

func GetNature(name string) (*Nature, error) {

	obj := &Nature{}
	err := callAPI("nature/", name, obj)
	return nil, err
}

func GetPokeathlonStat(name string) (*PokeathlonStat, error) {

	obj := &PokeathlonStat{}
	err := callAPI("pokeathlon-stat/", name, obj)
	return obj, err
}

//Pokemon returns a populated "Pokemon" struct
func GetPokemon(name string) (*Pokemon, error) {

	pokemon := &Pokemon{}
	err := callAPI("pokemon/", name, pokemon)
	return pokemon, err
}

func GetPokemonColor(name string) (*PokemonColor, error) {

	obj := &PokemonColor{}
	err := callAPI("pokemon-color/", name, obj)
	return obj, err
}

func GetPokemonForm(name string) (*PokemonForm, error) {

	obj := &PokemonForm{}
	err := callAPI("pokemon-form/", name, obj)
	return obj, err
}

func GetPokemonHabitat(name string) (*PokemonHabitat, error) {

	obj := &PokemonHabitat{}
	err := callAPI("pokemon-habitat/", name, obj)
	return obj, err
}

func GetPokemonShape(name string) (*PokemonShape, error) {

	obj := &PokemonShape{}
	err := callAPI("pokemon-shape/", name, obj)
	return obj, err
}

func GetPokemonSpecies(name string) (*PokemonSpecies, error) {

	obj := &PokemonSpecies{}
	err := callAPI("pokemon-species/", name, obj)
	return obj, err
}

func Getstat(name string) (*Stat, error) {

	obj := &Stat{}
	err := callAPI("stat/", name, obj)
	return obj, err
}

func GetType(name string) (*Type, error) {

	obj := &Type{}
	err := callAPI("type/", name, obj)
	return obj, err
}

func GetLanguage(name string) (*Language, error) {

	obj := &Language{}
	err := callAPI("language/", name, obj)
	return obj, err
}
