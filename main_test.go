package main

import (
	"log"
	"testing"
)

func TestGetVersionDefaultConf(t *testing.T) {
	tests := []struct {
		path  string
		regex string
	}{
		{"build.gradle", `build.gradle-regex`},
		{"", "missPath-regex"},
		{"pubspec.yaml", ""},
		{"", ""},
	}
	for _, test := range tests {
		switch true {
		case (test.path != "" && test.regex != ""): //all conf
			GetVersion(test.path, test.regex)
			log.Println("return vers, nil")
		case (test.path == "" && test.regex != ""): //no path
			for defPath := range defaultConfFiles {
				vers, err := GetVersion(defPath, test.regex)
				if err != nil {
					log.Printf("no path err: %s", err)
					continue
				}
				if vers == "" {
					log.Println("no searching vers")
				}
				log.Println("search vers: ", vers)
				break
			}
		case (test.path != "" && test.regex != ""): //no regex

		case (test.path != "" && test.regex != ""): //no conf
		}
		// for defPath, defReg := range defaultConfFiles {
		// 	log.Printf("path: %s,  req: %s", defPath, defReg)
		// }
	}
}
