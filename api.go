package pokego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GriMmzlyTM/pokego/types"
)

func buildURL(dir string, obj string) string {
	return fmt.Sprintf("%s%s%s", ops.URL, dir, strings.ToLower(obj))
}

//Calls the API. Used if there's no cache, or cache is expired.
func getAPI(dir string, name string) (*http.Response, error) {
	url := buildURL(dir, name)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

//callAPI is called by all the global functions to fetch data and return it.
//If Cache is set to true, it will attempt to get data from the cache, or create a new cache from data if one doesn't exist.
//If cache is set to false, it'll call the API directly and return the data
func callAPI(dir string, name string, data interface{}) error {

	if ops == (Options{}) {
		panic("Pokego Options not set! Please create a Options type variable, and asign it with func SetOptions()")
	}

	var content io.Reader

	if ops.Cache {
		cont, cacheErr := checkCache(dir, name)

		if cacheErr != nil {
			return cacheErr
		}

		content = cont
	} else {
		res, err := getAPI(dir, name)

		if err != nil {
			return err
		}

		content = res.Body
	}

	//Populates the struct
	jsonErr := json.NewDecoder(content).Decode(&data)

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
func BerryFlavor(name string) (*types.BerryFlavor, error) {

	obj := &types.BerryFlavor{}
	err := callAPI("berry-flavor/", name, obj)
	return obj, err
}

func ContestType(name string) (*types.ContestType, error) {

	obj := &types.ContestType{}
	err := callAPI("contest-type/", name, obj)
	return obj, err
}

func ContestEffect(id string) (*types.ContestEffect, error) {

	obj := &types.ContestEffect{}
	err := callAPI("contest-effect/", id, obj)
	return obj, err
}

func SuperContestEffect(id string) (*types.SuperContestEffect, error) {

	obj := &types.SuperContestEffect{}
	err := callAPI("super-contest-effect/", id, obj)
	return obj, err
}

func EncounterMethod(name string) (*types.EncounterMethod, error) {

	obj := &types.EncounterMethod{}
	err := callAPI("encounter-method/", name, obj)
	return obj, err
}

func EncounterCondition(name string) (*types.EncounterCondition, error) {

	obj := &types.EncounterCondition{}
	err := callAPI("encounter-condition/", name, obj)
	return obj, err
}

func EncounterConditionValue(name string) (*types.EncounterConditionValue, error) {

	obj := &types.EncounterConditionValue{}
	err := callAPI("encounter-condition-value/", name, obj)
	return obj, err
}

func EvolutionChain(name string) (*types.EvolutionChain, error) {

	obj := &types.EvolutionChain{}
	err := callAPI("evolution-chain/", name, obj)
	return obj, err
}

func EvolutionTrigger(name string) (*types.EvolutionTrigger, error) {

	obj := &types.EvolutionTrigger{}
	err := callAPI("evolution-trigger/", name, obj)
	return obj, err
}

func Generation(name string) (*types.Generation, error) {

	obj := &types.Generation{}
	err := callAPI("generation/", name, obj)
	return obj, err
}

func Pokedex(name string) (*types.Pokedex, error) {

	obj := &types.Pokedex{}
	err := callAPI("pokedex/", name, obj)
	return obj, err
}

func Version(name string) (*types.Version, error) {

	obj := &types.Version{}
	err := callAPI("version/", name, obj)
	return obj, err
}

func VersionGroup(name string) (*types.VersionGroup, error) {

	obj := &types.VersionGroup{}
	err := callAPI("version-group/", name, obj)
	return obj, err
}

func Item(name string) (*types.Item, error) {

	obj := &types.Item{}
	err := callAPI("item/", name, obj)
	return obj, err
}

func ItemAttribute(name string) (*types.ItemAttribute, error) {

	obj := &types.ItemAttribute{}
	err := callAPI("item-attribute/", name, obj)
	return obj, err
}

func ItemCategory(name string) (*types.ItemCategory, error) {

	obj := &types.ItemCategory{}
	err := callAPI("item-category/", name, obj)
	return obj, err
}

func ItemFlingEffect(name string) (*types.ItemFlingEffect, error) {

	obj := &types.ItemFlingEffect{}
	err := callAPI("item-fling-effect/", name, obj)
	return obj, err
}

func ItemPocket(name string) (*types.ItemPocket, error) {

	obj := &types.ItemPocket{}
	err := callAPI("item-pocket/", name, obj)
	return obj, err
}

func Location(name string) (*types.Location, error) {

	obj := &types.Location{}
	err := callAPI("location/", name, obj)
	return obj, err
}

func LocationArea(name string) (*types.LocationArea, error) {

	obj := &types.LocationArea{}
	err := callAPI("location-area/", name, obj)
	return obj, err
}

func PalParkArea(name string) (*types.PalParkArea, error) {

	obj := &types.PalParkArea{}
	err := callAPI("pal-park-area/", name, obj)
	return obj, err
}

func Region(name string) (*types.Region, error) {

	obj := &types.Region{}
	err := callAPI("region/", name, obj)
	return obj, err
}

func Machine(name string) (*types.Machine, error) {

	obj := &types.Machine{}
	err := callAPI("machine/", name, obj)
	return obj, err
}

func Move(name string) (*types.Move, error) {

	obj := &types.Move{}
	err := callAPI("move/", name, obj)
	return obj, err
}

func MoveAilment(name string) (*types.MoveAilment, error) {

	obj := &types.MoveAilment{}
	err := callAPI("move-ailment/", name, obj)
	return obj, err
}

func MoveBattleStyle(name string) (*types.MoveBattleStyle, error) {

	obj := &types.MoveBattleStyle{}
	err := callAPI("move-battle-style/", name, obj)
	return obj, err
}

func MoveCategory(name string) (*types.MoveCategory, error) {

	obj := &types.MoveCategory{}
	err := callAPI("move-category/", name, obj)
	return obj, err
}

func MoveDamageClass(name string) (*types.MoveDamageClass, error) {

	obj := &types.MoveDamageClass{}
	err := callAPI("move-damage-class/", name, obj)
	return obj, err
}

func MoveLearnMethod(name string) (*types.MoveLearnMethod, error) {

	obj := &types.MoveLearnMethod{}
	err := callAPI("move-learn-method/", name, obj)
	return obj, err
}

func MoveTarget(name string) (*types.MoveTarget, error) {

	obj := &types.MoveTarget{}
	err := callAPI("move-target/", name, obj)
	return obj, err
}

func Ability(name string) (*types.Ability, error) {

	obj := &types.Ability{}
	err := callAPI("ability/", name, obj)
	return obj, err
}

func Characteristic(name string) (*types.Characteristic, error) {

	obj := &types.Characteristic{}
	err := callAPI("characteristic/", name, obj)
	return obj, err
}

func EggGroup(name string) (*types.EggGroup, error) {

	obj := &types.EggGroup{}
	err := callAPI("egg-group/", name, obj)
	return obj, err
}

func Gender(name string) (*types.Gender, error) {

	obj := &types.Gender{}
	err := callAPI("gender/", name, obj)
	return obj, err
}

func GrowthRate(name string) (*types.GrowthRate, error) {

	obj := &types.GrowthRate{}
	err := callAPI("growth-rate/", name, obj)
	return obj, err
}

func Nature(name string) (*types.Nature, error) {

	obj := &types.Nature{}
	err := callAPI("nature/", name, obj)
	return nil, err
}

func PokeathlonStat(name string) (*types.PokeathlonStat, error) {

	obj := &types.PokeathlonStat{}
	err := callAPI("pokeathlon-stat/", name, obj)
	return obj, err
}

//Pokemon returns a populated "Pokemon" struct
func Pokemon(name string) (*types.Pokemon, error) {

	pokemon := &types.Pokemon{}
	err := callAPI("pokemon/", name, pokemon)
	return pokemon, err
}

func PokemonColor(name string) (*types.PokemonColor, error) {

	obj := &types.PokemonColor{}
	err := callAPI("pokemon-color/", name, obj)
	return obj, err
}

func PokemonForm(name string) (*types.PokemonForm, error) {

	obj := &types.PokemonForm{}
	err := callAPI("pokemon-form/", name, obj)
	return obj, err
}

func PokemonHabitat(name string) (*types.PokemonHabitat, error) {

	obj := &types.PokemonHabitat{}
	err := callAPI("pokemon-habitat/", name, obj)
	return obj, err
}

func PokemonShape(name string) (*types.PokemonShape, error) {

	obj := &types.PokemonShape{}
	err := callAPI("pokemon-shape/", name, obj)
	return obj, err
}

func PokemonSpecies(name string) (*types.PokemonSpecies, error) {

	obj := &types.PokemonSpecies{}
	err := callAPI("pokemon-species/", name, obj)
	return obj, err
}

func stat(name string) (*types.Stat, error) {

	obj := &types.Stat{}
	err := callAPI("stat/", name, obj)
	return obj, err
}

func Type(name string) (*types.Type, error) {

	obj := &types.Type{}
	err := callAPI("type/", name, obj)
	return obj, err
}

func Language(name string) (*types.Language, error) {

	obj := &types.Language{}
	err := callAPI("language/", name, obj)
	return obj, err
}
