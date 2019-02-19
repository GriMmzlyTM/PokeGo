package test

import (
	"os"
	"testing"

	"github.com/GriMmzlyTM/pokego"
)

func TestMain(m *testing.M) {

	ops := pokego.Options{
		URL:      "https://pokeapi.co/api/v2/",
		Cache:    true,
		CacheDir: `.\testcache\`,
	}

	pokego.SetOptions(ops)

	os.Exit(m.Run())
}

func TestPokemon(t *testing.T) {
	poke, pokeErr := pokego.GetPokemon("Pikachu")

	if pokeErr != nil {
		t.Error("pokego.Pokemon(\"Pikachu\") returned an error. Check options URL or JSON conversion.")
	}
	if poke.BaseExperience == 0 && poke.Moves[0].Move.Name == "" {
		t.Error("Pokemon was improperly set, but no error was thrown.")
	}
}

func TestMove(t *testing.T) {
	move, moveErr := pokego.GetMove("hyper-BeAm")

	if moveErr != nil {
		t.Error("pokego.Move(\"hyper-BeAm\") returned an error. Check options URL or JSON conversion.")
	}
	if move.Accuracy == 0 && move.Name == "" {
		t.Error("Move not properly set, but no error was thrown")
	}

}

func TestVersion(t *testing.T) {

	_, versionErr := pokego.GetVersion("silver")

	if versionErr != nil {
		t.Error("Could not get version. Check options or Json")
	}

}

func TestBerry(t *testing.T) {
	berry, _ := pokego.GetBerry("hondew")

	t.Log(berry.Name)

}
