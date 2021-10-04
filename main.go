package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {

	var str1 struct {
		name string
		id   int
	}

	var str2 struct {
		name string
		id   int
	}

	str1.name = "vadim"
	str2.id = 1

	str2.name = ""
	str2.id = 2

	var names = &str1.name
	str2.name = *names
	fmt.Println("str1: ", str1.name, "    str2: ", str2.name, "    names: ", *names)
	str1.name = ""
	fmt.Println("str1: ", str1.name, "    str2: ", str2.name, "    names: ", *names)
	var str1_name = &str1.name
	str1.name = ""
	fmt.Println(*str1_name)
	fmt.Println(str1.name)

	str1.name = "str1 name"
	str1.id = 1

	str2.name = str1.name

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
