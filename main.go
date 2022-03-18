package main

import (
	"errors"
	"log"
	"regexp"

	"github.com/hashicorp/go-version"
)

func main() {
	version.NewVersion("1.0.1")
}

var UnmarshalBuildGradle = func(content []byte) error {
	regex, _ := regexp.Compile(`versionName "([^\t\n\f\r]+)"`)
	res := regex.FindStringSubmatch(string(content))
	if len(res) == 0 {
		return errors.New("err")
	}
	vers := res[1]
	log.Println(vers)

	return nil
}
