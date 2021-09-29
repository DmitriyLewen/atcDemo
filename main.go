package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("grable.properties")
	if err != nil {
		fmt.Println("error read file: ", err)
	}
	s := string(bytes)
	vers := findnpmrc(s)
	templateVers := "v{{.version}}"
	result := strings.Replace(templateVers, ".version", vers, -1)
	fmt.Println(result)
}

func findnpmrc(fileString string) string {
	regex, _ := regexp.Compile("version=([^\t\n\f\r]*)")
	res := regex.FindStringSubmatch(fileString)
	if len(res) != 0 {
		return res[1]
	}
	return ""
}
