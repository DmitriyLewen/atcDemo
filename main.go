package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	resp := "ewogICAgIm5hbWUiOiAibXlfcGFja2FnZSIsCiAgICAiZGVzY3JpcHRpb24i\nOiAiIiwKICAgICJ2ZXJzaW9uIjogIjEuMC4yIiwKICAgICJzY3JpcHRzIjog\newogICAgICAidGVzdCI6ICJlY2hvIFwiRXJyb3I6IG5vIHRlc3Qgc3BlY2lm\naWVkXCIgJiYgZXhpdCAxIgogICAgfSwKICAgICJyZXBvc2l0b3J5Ijogewog\nICAgICAidHlwZSI6ICJnaXQiLAogICAgICAidXJsIjogImh0dHBzOi8vZ2l0\naHViLmNvbS9tb25hdGhlb2N0b2NhdC9teV9wYWNrYWdlLmdpdCIKICAgIH0s\nCiAgICAia2V5d29yZHMiOiBbXSwKICAgICJhdXRob3IiOiAiIiwKICAgICJs\naWNlbnNlIjogIklTQyIsCiAgICAiYnVncyI6IHsKICAgICAgInVybCI6ICJo\ndHRwczovL2dpdGh1Yi5jb20vbW9uYXRoZW9jdG9jYXQvbXlfcGFja2FnZS9p\nc3N1ZXMiCiAgICB9LAogICAgImhvbWVwYWdlIjogImh0dHBzOi8vZ2l0aHVi\nLmNvbS9tb25hdGhlb2N0b2NhdC9teV9wYWNrYWdlIgogIH0=\n"

	respStr, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		log.Printf("decode err : %v", err)
	}
	fmt.Printf("decode: %s", respStr)
	fmt.Println("+++++")
	jsonT()

	str1 := "Cjxwcm9qZWN0IHhtbG5zPSJodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy9QT00vNC4wLjAiIHhtbG5zOnhzaT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiCgl4c2k6c2NoZW1hTG9jYXRpb249Imh0dHA6Ly9tYXZlbi5hcGFjaGUub3JnL1BPTS80LjAuMCBodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy94c2QvbWF2ZW4tNC4wLjAueHNkIj4KCTx2ZXJzaW9uPjU8L3ZlcnNpb24"
	str2 := "Cjxwcm9qZWN0IHhtbG5zPSJodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy9QT00vNC4wLjAiIHhtbG5zOnhzaT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiCgl4c2k6c2NoZW1hTG9jYXRpb249Imh0dHA6Ly9tYXZlbi5hcGFjaGUub3JnL1BPTS80LjAuMCBodHRwOi8vbWF2ZW4uYXBhY2hlLm9yZy94c2QvbWF2ZW4tNC4wLjAueHNkIj4KCTx2ZXJzaW9uPjU8L3ZlcnNpb24"
	fmt.Println("compare string :", str1 == str2)

	url1 := "https://api.github.com/repos/Codertocat/Hello-World/contents/contents/pom.xml"
	url2 := "https://api.github.com/repos/Codertocat/Hello-World/contents/contents/pom.xml?ref=6113728f27ae82c7b1a177c8d03f9e96e0adf246"
	fmt.Println("regexghservise url1: ", regexghserviseOld(url1))
	fmt.Println("regexghservise url2: ", regexghserviseOld(url2))
	fmt.Println("regexghservise url1: ", regexghserviseNew(url1))
	fmt.Println("regexghservise url2: ", regexghserviseNew(url2))

	strm := madeСaptionToTemplate(`{{.version}}Vtee`, "5")
	fmt.Println("strm: ", strm)

	regStr := "([^\t\n\f\r]*?)$"
	fmt.Println(" path " + regStr)

	fmt.Println("contents/pom.xml: ", filepath.Base(`project/contents/pom.xml`))

	fmt.Println("contents/pom.xml/: ", filepath.Base("contents//pom.xml/"))

	text := `
## gradle.properties

# Common Android settings
android.compileSdkVersion=28
android.applicationId=com.example
android.targetSdkVersion=28
android.minSdkVersion=21
android.version=2
android.versionName=1.2
plugin.com.github.ben-manes.versions=0.25.0
plugin.de.fayard.buildSrcVersions=0.6.1
version.com.android.tools.build..gradle=3.5.0
version.play-services-location=17.0.0
version.bottom-navigation-bar=2.1.0
version.lifecycle-extensions=2.0.0
#                # available=2.1.0
version.org.jetbrains.kotlin=1.3.31
#                # available=1.3.50
version.appcompat=1.1.0-rc01
#     # available=1.1.0
version.cardview=1.0.0
version.core-ktx=1.0.2
#    # available=1.1.0
# ....
version=1.0.34
org.gradle.caching=true
org.gradle.parallel=true
org.gradle.caching.debug=false
org.gradle.configureondemand=false
org.gradle.daemon.idletimeout= 10800000
org.gradle.console=auto
#org.gradle.java.home=(path to JDK home)
#org.gradle.warning.mode=(all,none,summary)
#org.gradle.workers.max=(max # of worker processes)
# org.gradle.priority=(low,normal)
version=1.0.35`

	findVers(text)

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

func findVers(text string) string {
	//regex, _ := regexp.Compile(`^version=([^\t\n\f\r]*)|[\t\n\f\r]version=([^\t\n\f\r]*)`)
	regex, _ := regexp.Compile(`^[\t\n\f\r]*version=([^\t\n\f\r]*)`)
	res := regex.FindStringSubmatch(text)
	for i, r := range res {
		fmt.Printf("res%d: %s\n", i, r)
	}
	return res[1]
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

type jsonStruct struct {
	ar string
}

func jsonT() {
	fmt.Println("____jsonT start_____")
	str := []jsonStruct{
		{"one"},
		{"two"},
	}
	data, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	file, err := os.OpenFile("jsonfile.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	fileRead, err := os.ReadFile("jsonfile.txt")
	if err != nil {
		log.Println(err)
	}
	var titles []struct{ Title string }
	err = json.Unmarshal(fileRead, &titles)
	if err != nil {
		log.Println(err)
	}
	file.WriteString(string(data))
	fmt.Println("____jsonT end_____")
}
