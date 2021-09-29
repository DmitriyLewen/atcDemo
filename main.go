package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	fileString := "npm config set init.author.name \"Hiro Protagonist\" \nnpm config set init.author.email \"hiro@showcrash.io\" \nnpm config set init.author.url \"http://hiro.snowcrash.io\" \nnpm config set init.license \"MIT\" \nnpm config set init.version \"0.0.1\""
	findnpmrc(fileString)
	bytes, err := ioutil.ReadFile(".npmrc")
	s := string(bytes)
	if err != nil {
		fmt.Println("error read file: ", err)
	}
	fmt.Println("1:", s)
}

func findnpmrc(fileString string) {
	regex, _ := regexp.Compile("npm config set init.version \"([^\t\n\f\r]*?)\"")
	res := regex.FindStringSubmatch(fileString)
	fmt.Println(res[1])
}
