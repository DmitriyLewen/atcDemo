package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	str1 := "Cjxwcm9qZWN0IHhtbG5zPSJodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy9QT00vNC4wLjAiIHhtbG5zOnhzaT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiCgl4c2k6c2NoZW1hTG9jYXRpb249Imh0dHA6Ly9tYXZlbi5hcGFjaGUub3JnL1BPTS80LjAuMCBodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy94c2QvbWF2ZW4tNC4wLjAueHNkIj4KCTx2ZXJzaW9uPjU8L3ZlcnNpb24"
	str2 := "Cjxwcm9qZWN0IHhtbG5zPSJodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy9QT00vNC4wLjAiIHhtbG5zOnhzaT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiCgl4c2k6c2NoZW1hTG9jYXRpb249Imh0dHA6Ly9tYXZlbi5hcGFjaGUub3JnL1BPTS80LjAuMCBodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy94c2QvbWF2ZW4tNC4wLjAueHNkIj4KCTx2ZXJzaW9uPjU8L3ZlcnNpb24"
	fmt.Println("compare string :", str1 == str2)

	url1 := "https://api.github.com/repos/Codertocat/Hello-World/contents/contents/pom.xml"
	url2 := "https://api.github.com/repos/Codertocat/Hello-World/contents/contents/pom.xml?ref=6113728f27ae82c7b1a177c8d03f9e96e0adf246"
	fmt.Println("regexghservise url1: ", regexghserviseOld(url1))
	fmt.Println("regexghservise url2: ", regexghserviseOld(url2))
	fmt.Println("regexghservise url1: ", regexghserviseNew(url1))
	fmt.Println("regexghservise url2: ", regexghserviseNew(url2))

	strm := madeСaptionToTemplate(`{{.version}}Vte`, "5")
	fmt.Println("strm: ", strm)

	regStr := "([^\t\n\f\r]*?)$"
	fmt.Println("path " + regStr)

	fmt.Println("contents/pom.xml : ", filepath.Base(`project////contents/pom.xml`))

	fmt.Println("contents/pom.xml/ : ", filepath.Base("contents/pom.xml/"))
	// if fetcher == nil {
	// 	log.Printf("Error: non support Path = %q", settings.Path)
	// 	return
	// }

	// regStr := "([^\t\n\f\r]*)"
	// regexPath, err := regexp.Compile("path: " + regStr)
	// if err != nil {
	// 	return err
	// }
	// regexTemplate, _ := regexp.Compile("template: " + regStr)
	// if err != nil {
	// 	return err
	// }
	// regexBehavior, _ := regexp.Compile("behavior: " + regStr)
	// if err != nil {
	// 	return err
	// }
	// path := regexPath.FindStringSubmatch(string(content))
	// template := regexTemplate.FindStringSubmatch(string(content))
	// behavior := regexBehavior.FindStringSubmatch(string(content))
	// if len(path) > 0 {
	// 	atcSettingsPtr.Path = path[1]
	// }
	// if len(template) > 0 {
	// 	atcSettingsPtr.Template = template[1]
	// }
	// if len(behavior) > 0 {
	// 	atcSettingsPtr.Behavior = behavior[1]
	// }
	// //err := yaml.Unmarshal([]byte(content), atcSettingsPtr)
	// return nil

	// {`template: v{{.version}}`, `Added a new version for "Codertocat/Hello-World": "v5"`, `v5`},
	// 	{`template: vTest{{.version}}`, `Added a new version for "Codertocat/Hello-World": "vTest5"`, `vTest5`},
	// 	{`template: {{.version}}Vte`, `Added a new version for "Codertocat/Hello-World": "5Vte"`, `5Vte`},
	// 	{`template: vVv{.version}`, `Added a new version for "Codertocat/Hello-World": "v5"`, `v5`},
	// 	{`template: `, `Added a new version for "Codertocat/Hello-World": "v5"`, `v5`},
	// 	{`template: v{{.version}}`, `Added a new version for "Codertocat/Hello-World": "v5"`, `v5`},
	// 	{`template: vVv{{.verson}}`, `Added a new version for "Codertocat/Hello-World": "v5"`, `v5`},
	// 	{`template: vvV{{version}}`, `Added a new version for "Codertocat/Hello-World": "v5"`, `v5`},
}

func madeСaptionToTemplate(template, version string) string {
	if !strings.Contains(template, "{{.version}}") {
		return "v" + version
	}
	return strings.Replace(template, "{{.version}}", version, -1)
}

func regexghserviseOld(url string) bool {
	matched, err := regexp.MatchString("pom\\.xml\\?ref=", url)
	if err != nil {
		return false
	}
	return matched
}

func regexghserviseNew(url string) bool {
	matched, err := regexp.MatchString("pom\\.xml$", url)
	if err != nil {
		return false
	}
	return matched
}
func findnpmrc(fileString string) string {
	reg := `contents/pom.xml|gradle.properties|.npmrc`
	regex, _ := regexp.Compile(reg)
	str1 := `gradle.properties`
	str2 := `contents/pom.xml`
	matched, _ := regexp.MatchString(reg, str1)
	fmt.Println(matched)
	matched, _ = regexp.MatchString(reg, str2)
	fmt.Println(matched)
	res := regex.FindStringSubmatch(fileString)
	if len(res) != 0 {
		return res[1]
	}
	return ""
}
