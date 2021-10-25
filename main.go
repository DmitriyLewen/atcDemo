package main

import (
	"errors"
	"log"
	"regexp"
)

func main() {

}

var unmarshalBuildGradle = func(content []byte) error {
	regex, _ := regexp.Compile(`versionName "([^\t\n\f\r]+)"`)
	res := regex.FindStringSubmatch(string(content))
	if len(res) == 0 {
		return errors.New("err")
	}
	vers := res[1]
	log.Println(vers)
	return nil
}
