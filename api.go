package pokego

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pokego/pokego/types"
)

func buildURL(dir string, obj string) string {
	return ops.URL + dir + obj
}

func callAPI(dir string, name string, data interface{}) error {

	if ops == (Options{}) {
		panic("Pokego Options not set! Please create a Options type variable, and asign it with func SetOptions()")
	}

	url := buildURL(dir, name)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonErr := json.NewDecoder(res.Body).Decode(&data)

	return jsonErr
}

//Berry returns a pointer a Berry type var
func Berry(name string) (*types.Berry, error) {
	berry := &types.Berry{}
	err := callAPI("berry/", name, berry)
	return berry, err
}

//BerryFirmness returns a pointer to a BerryFirmness var
func BerryFirmness(name string) (*types.BerryFirmness, error) {

	obj := &types.BerryFirmness{}
	err := callAPI("berry-firmness/", name, obj)
	return obj, err
}

//BerryFlavor returns a pointer.
func BerryFlavor(name string) *types.BerryFlavor {

	obj := &types.BerryFlavor{}
	callAPI("berry-flavor/", name, obj)
	return obj
}

func ContestType(name string) *types.ContestType {

	obj := &types.ContestType{}
	callAPI("contest-type/", name, obj)
	return obj
}

func ContestEffect(id string) *types.ContestEffect {

	obj := &types.ContestEffect{}
	callAPI("contest-effect/", id, obj)
	return obj
}

func SuperContestEffect(id string) *types.SuperContestEffect {

	obj := &types.SuperContestEffect{}
	callAPI("super-contest-effect/", id, obj)
	return obj
}

func EncounterMethod(name string) *types.EncounterMethod {

	obj := &types.EncounterMethod{}
	callAPI("encounter-method/", name, obj)
	return obj
}

func EncounterCondition(name string) *types.EncounterCondition {

	obj := &types.EncounterCondition{}
	callAPI("encounter-condition/", name, obj)
	return obj
}

func EncounterConditionValue(name string) *types.EncounterConditionValue {

	obj := &types.EncounterConditionValue{}
	callAPI("encounter-condition-value/", name, obj)
	return obj
}

func EvolutionChain(name string) *types.EvolutionChain {

	obj := &types.EvolutionChain{}
	callAPI("evolution-chain/", name, obj)
	return obj
}

func EvolutionTrigger(name string) *types.EvolutionTrigger {

	obj := &types.EvolutionTrigger{}
	callAPI("evolution-trigger/", name, obj)
	return obj
}

func Generation(name string) *types.Generation {

	obj := &types.Generation{}
	callAPI("generation/", name, obj)
	return obj
}

func Pokedex(name string) *types.Pokedex {

	obj := &types.Pokedex{}
	callAPI("pokedex/", name, obj)
	return obj
}

func Version(name string) *types.Version {

	obj := &types.Version{}
	callAPI("version/", name, obj)
	return obj
}

func VersionGroup(name string) *types.VersionGroup {

	obj := &types.VersionGroup{}
	callAPI("version-group/", name, obj)
	return obj
}

func Item(name string) *types.Item {

	obj := &types.Item{}
	callAPI("item/", name, obj)
	return obj
}

func ItemAttribute(name string) *types.ItemAttribute {

	obj := &types.ItemAttribute{}
	callAPI("item-attribute/", name, obj)
	return obj
}

func ItemCategory(name string) *types.ItemCategory {

	obj := &types.ItemCategory{}
	callAPI("item-category/", name, obj)
	return obj
}

func ItemFlingEffect(name string) *types.ItemFlingEffect {

	obj := &types.ItemFlingEffect{}
	callAPI("item-fling-effect/", name, obj)
	return obj
}

func ItemPocket(name string) *types.ItemPocket {

	obj := &types.ItemPocket{}
	callAPI("item-pocket/", name, obj)
	return obj
}

func Location(name string) *types.Location {

	obj := &types.Location{}
	callAPI("location/", name, obj)
	return obj
}

func LocationArea(name string) *types.LocationArea {

	obj := &types.LocationArea{}
	callAPI("location-area/", name, obj)
	return obj
}

func PalParkArea(name string) *types.PalParkArea {

	obj := &types.PalParkArea{}
	callAPI("pal-park-area/", name, obj)
	return obj
}

func Region(name string) *types.Region {

	obj := &types.Region{}
	callAPI("region/", name, obj)
	return obj
}

func Machine(name string) *types.Machine {

	obj := &types.Machine{}
	callAPI("machine/", name, obj)
	return obj
}

func Move(name string) *types.Move {

	obj := &types.Move{}
	callAPI("move/", name, obj)
	return obj
}

func MoveAilment(name string) *types.MoveAilment {

	obj := &types.MoveAilment{}
	callAPI("move-ailment/", name, obj)
	return obj
}

func MoveBattleStyle(name string) *types.MoveBattleStyle {

	obj := &types.MoveBattleStyle{}
	callAPI("move-battle-style/", name, obj)
	return obj
}

func MoveCategory(name string) *types.MoveCategory {

	obj := &types.MoveCategory{}
	callAPI("move-category/", name, obj)
	return obj
}

func MoveDamageClass(name string) *types.MoveDamageClass {

	obj := &types.MoveDamageClass{}
	callAPI("move-damage-class/", name, obj)
	return obj
}

func MoveLearnMethod(name string) *types.MoveLearnMethod {

	obj := &types.MoveLearnMethod{}
	callAPI("move-learn-method/", name, obj)
	return obj
}

func MoveTarget(name string) *types.MoveTarget {

	obj := &types.MoveTarget{}
	callAPI("move-target/", name, obj)
	return obj
}

func Ability(name string) *types.Ability {

	obj := &types.Ability{}
	callAPI("ability/", name, obj)
	return obj
}

func Characteristic(name string) *types.Characteristic {

	obj := &types.Characteristic{}
	callAPI("characteristic/", name, obj)
	return obj
}

func EggGroup(name string) *types.EggGroup {

	obj := &types.EggGroup{}
	callAPI("egg-group/", name, obj)
	return obj
}

func Gender(name string) *types.Gender {

	obj := &types.Gender{}
	callAPI("gender/", name, obj)
	return obj
}

func GrowthRate(name string) *types.GrowthRate {

	obj := &types.GrowthRate{}
	callAPI("growth-rate/", name, obj)
	return obj
}

func Nature(name string) *types.Nature {

	obj := &types.Nature{}
	callAPI("nature/", name, obj)
	return obj
}

func PokeathlonStat(name string) *types.PokeathlonStat {

	obj := &types.PokeathlonStat{}
	callAPI("pokeathlon-stat/", name, obj)
	return obj
}

//Pokemon returns a populated "Pokemon" struct
func Pokemon(name string) *types.Pokemon {

	pokemon := &types.Pokemon{}
	callAPI("pokemon/", name, pokemon)
	return pokemon
}

func PokemonColor(name string) *types.PokemonColor {

	obj := &types.PokemonColor{}
	callAPI("pokemon-color/", name, obj)
	return obj
}

func PokemonForm(name string) *types.PokemonForm {

	obj := &types.PokemonForm{}
	callAPI("pokemon-form/", name, obj)
	return obj
}

func PokemonHabitat(name string) *types.PokemonHabitat {

	obj := &types.PokemonHabitat{}
	callAPI("pokemon-habitat/", name, obj)
	return obj
}

func PokemonShape(name string) *types.PokemonShape {

	obj := &types.PokemonShape{}
	callAPI("pokemon-shape/", name, obj)
	return obj
}

func PokemonSpecies(name string) *types.PokemonSpecies {

	obj := &types.PokemonSpecies{}
	callAPI("pokemon-species/", name, obj)
	return obj
}

func stat(name string) *types.Stat {

	obj := &types.Stat{}
	callAPI("stat/", name, obj)
	return obj
}

func Type(name string) *types.Type {

	obj := &types.Type{}
	callAPI("type/", name, obj)
	return obj
}

func Language(name string) *types.Language {

	obj := &types.Language{}
	callAPI("language/", name, obj)
	return obj
}
