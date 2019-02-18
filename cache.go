package pokego

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//Check if the json file exists in the cache
//If it exist, it returns the JSON string. Else, it returns an empty string.
func checkCache(dir string, name string) (io.Reader, error) {

	//EXAMPLE: CACHEDIR/pokemon/pikachu.json
	fileDir := fmt.Sprintf("%s%s%s.json", ops.CacheDir, dir, name)

	file, openErr := os.Open(fileDir)
	defer file.Close()
	if openErr != nil {
		if os.IsNotExist(openErr) {
			content, cacheErr := createCache(dir, name)

			return content, cacheErr
		}
		return nil, openErr
	}

	jsonByte, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return nil, readErr
	}

	jsonString := string(jsonByte[:])

	return strings.NewReader(jsonString), nil
}

func createCache(dir string, name string) (io.Reader, error) {

	res, resErr := getAPI(dir, name)

	if resErr != nil {
		return nil, resErr
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	fileDir := fmt.Sprintf("%s%s%s.json", ops.CacheDir, dir, strings.ToLower(name))
	if _, err := os.Stat(fmt.Sprintf("%s%s", ops.CacheDir, dir)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("%s%s", ops.CacheDir, dir), 0666)
	}

	writeErr := ioutil.WriteFile(fileDir, body, 0666)

	content := bytes.NewReader(body)

	return content, writeErr
}
