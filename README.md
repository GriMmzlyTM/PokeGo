# PokéGO

[![Go Report Card](https://goreportcard.com/badge/github.com/GriMmzlyTM/pokego)](https://goreportcard.com/report/github.com/GriMmzlyTM/pokego)  [![godoc](https://img.shields.io/badge/godoc-pokego-blue.svg)](https://godoc.org/github.com/GriMmzlyTM/pokego)  ![GitHub last commit](https://img.shields.io/github/last-commit/GriMmzlyTM/pokego.svg?colorB=blue&style=flat)  ![GitHub](https://img.shields.io/github/license/GriMmzlyTM/pokego.svg?colorB=blue&style=flat)

Pokego is a golang wrapper for the pokeapi RESTful API.

The JSON returned from the API is converted to a struct. Any data type available in the API has been replicated as a struct here. [Read the pokeapi documentation for more info.](https://pokeapi.co/docs/v2.html)

Pokego is still a work in progress.


## Getting Started

Open a terminal and download the library
```bash
go get github.com/GriMmzlyTM/pokego
```

Once the library has been downloaded, import the library
```go
import (
    pokego "github.com/GriMmzlyTM/pokego"
)
```

## How to

Pokego is meant to be as easy to use as possible. Once imported, you'll need to set the options:

```go
//Make a new Options variable
ops := pokego.Options {
            URL:      "https://pokeapi.co/api/v2/",
            Cache:    true,
            CacheDir: `Dir/To/Cache/Folder`,
            CacheExpiryDays: 20,
    }
    
//Set the options
pokego.SetOptions(ops)
```

Once your options are set, you can start calling the API

```go
//This will create a "Pokemon" struct, and populate it with the charizard data
charizard, err := pokego.GetPokemon("charizard")

if err != nil {
    panic("Something went wrong!!")
}

//This will create a new variable of type Move, and populate it with hyper-beam data. 
var coolMove Move 
moveErr := coolMove.Get("hyper-beam")

if moveErr != nil {
    panic("What's going on?!?!")
}

/*Both methods are viable, and which you choose to use depends on how you plan on using the data. 
Most structs contain a Get method, which allows you to override the data.*/ 
```

Once you have populated your struct, you can access the data as you normally would

```go
fmt.Println(charizard.Abilities[0])
```

For a full list of structs and methods, check the godoc badge and pokeapi documentation. All data structures that exist in the original API exist here as well. Strings are not case sensitive.

## Running the tests

To run the tests, open a terminal and cd into the "github.com/GriMmzlyTM/pokego/test" folder. Execute the go test command from there. 
```bash
go test -v
```

This will execute all tests, as well as populate the testcache to ensure caching is working. 

## Authors

* **Lorenzo Torelli** - *Dev* - [GriMmzlyTM](https://github.com/GriMmzlyTM)

See also the list of [contributors](https://github.com/GriMmzlyTM/pokego/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Pokeapi.co for creating the RESTful API

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)   ![Pokego badge](https://img.shields.io/badge/Pok%C3%A9-GO-blue.svg)
