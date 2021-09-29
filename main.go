package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	bytes, err := ioutil.ReadFile("grable.properties")
	if err != nil {
		fmt.Println("error read file: ", err)
	}
	s := string(bytes)
	findnpmrc(s)
}

func findnpmrc(fileString string) {
	regex, _ := regexp.Compile("version=([^\t\n\f\r]*)")
	res := regex.FindStringSubmatch(fileString)
	fmt.Println(res)
	if len(res) != 0 {
		fmt.Println("____________")
		fmt.Println(res[1])
	}

}
