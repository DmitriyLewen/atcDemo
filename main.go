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
	v := "1.0"
	result := strings.Replace(templateVers, "{{.version}}", vers, -1)
	fmt.Println(result)
	fmt.Println("v" + v)

	var mape = map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}

	fmt.Println(mape["1"])
	fmt.Println(mape["5"])
}

func findnpmrc(fileString string) string {
	regex, _ := regexp.Compile("version=([^\t\n\f\r]*)")
	res := regex.FindStringSubmatch(fileString)
	if len(res) != 0 {
		return res[1]
	}
	return ""
}
