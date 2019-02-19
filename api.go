package pokego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
		panic("Pokego Options not set! Please create a Options type variable, and assign it with func SetOptions()")
	}

	var content io.Reader

	if ops.Cache {
		cont, cacheErr := getCache(dir, name)
		
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
