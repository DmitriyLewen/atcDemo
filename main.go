package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"gopkg.in/yaml.v2"
)

var errUnmarshal = errors.New("unmarshal err")

var defaultConfFiles = map[string]string{
	"app/build.gradle": `versionName "([^\t\n\f\r]+)"`,
	"package.json":     `"version": "([^\t\n\f\r]+)",`,
	"pubspec.yaml":     `version: ([^\t\n\f\r]+)`,
	"contents/pom.xml": `<modelVersion>(.+)</modelVersion>`, //mb add to regex?????
}

func main() {
	// for defPath, defReg := range defaultConfFiles {
	// 	log.Printf("path: %s,  req: %s", defPath, defReg)
	// }
	content := `
	plugins {
		id 'com.android.application'
		id 'kotlin-android'
		id 'kotlin-kapt'
		id 'dagger.hilt.android.plugin'
	}
	
	android {
		compileSdk 31
	
		defaultConfig {
			applicationId "com.plcoding.cryptocurrencyappyt"
			minSdk 21
			targetSdk 31
			versionCode 1
			versionName "1.1"
	
			testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"
			vectorDrawables {
				useSupportLibrary true
			}
		}
	
		buildTypes {
			release {
				minifyEnabled false
				proguardFiles getDefaultProguardFile('proguard-android-optimize.txt'), 'proguard-rules.pro'
			}
		}
		compileOptions {
			sourceCompatibility JavaVersion.VERSION_1_8
			targetCompatibility JavaVersion.VERSION_1_8
		}
		kotlinOptions {
			jvmTarget = '1.8'
			useIR = true
		}
		buildFeatures {
			compose true
		}
		composeOptions {
			kotlinCompilerExtensionVersion compose_version
			kotlinCompilerVersion '1.5.21'
		}
		packagingOptions {
			resources {
				excludes += '/META-INF/{AL2.0,LGPL2.1}'
			}
		}
	}
	
	dependencies {
	
		implementation 'androidx.core:core-ktx:1.6.0'
		implementation 'androidx.appcompat:appcompat:1.3.1'
		implementation 'com.google.android.material:material:1.4.0'
		implementation "androidx.compose.ui:ui:$compose_version"
		implementation "androidx.compose.material:material:$compose_version"
		implementation "androidx.compose.ui:ui-tooling-preview:$compose_version"
		implementation 'androidx.lifecycle:lifecycle-runtime-ktx:2.3.1'
		implementation 'androidx.activity:activity-compose:1.3.1'
		testImplementation 'junit:junit:4.+'
		androidTestImplementation 'androidx.test.ext:junit:1.1.3'
		androidTestImplementation 'androidx.test.espresso:espresso-core:3.4.0'
		androidTestImplementation "androidx.compose.ui:ui-test-junit4:$compose_version"
		debugImplementation "androidx.compose.ui:ui-tooling:$compose_version"
	
		// Compose dependencies
		implementation "androidx.lifecycle:lifecycle-viewmodel-compose:1.0.0-alpha07"
		implementation "androidx.navigation:navigation-compose:2.4.0-alpha08"
		implementation "com.google.accompanist:accompanist-flowlayout:0.17.0"
	
		// Coroutines
		implementation 'org.jetbrains.kotlinx:kotlinx-coroutines-core:1.5.1'
		implementation 'org.jetbrains.kotlinx:kotlinx-coroutines-android:1.5.1'
	
		// Coroutine Lifecycle Scopes
		implementation "androidx.lifecycle:lifecycle-viewmodel-ktx:2.3.1"
		implementation "androidx.lifecycle:lifecycle-runtime-ktx:2.3.1"
	
		//Dagger - Hilt
		implementation "com.google.dagger:hilt-android:2.38.1"
		kapt "com.google.dagger:hilt-android-compiler:2.37"
		implementation "androidx.hilt:hilt-lifecycle-viewmodel:1.0.0-alpha03"
		kapt "androidx.hilt:hilt-compiler:1.0.0"
		implementation 'androidx.hilt:hilt-navigation-compose:1.0.0-alpha03'
	
		// Retrofit
		implementation 'com.squareup.retrofit2:retrofit:2.9.0'
		implementation 'com.squareup.retrofit2:converter-gson:2.9.0'
		implementation "com.squareup.okhttp3:okhttp:5.0.0-alpha.2"
		implementation "com.squareup.okhttp3:logging-interceptor:5.0.0-alpha.2"
	}`
	regex, _ := regexp.Compile(`versionNameaaaa ".+"`)
	res := regex.FindString(string(content))
	fmt.Printf("find string:%s", res)
}

func GetVersion(path, regex string) (string, error) {
	log.Printf("GetVersion in, path: %s, regex: %s", path, regex)
	if path != "contents/pom.xml" {
		return "", errUnmarshal
	}
	return "1.0", nil
}

type AtcSettings struct {
	Path          string `yaml:"path"`
	Behavior      string `yaml:"behavior"`
	Template      string `yaml:"template"`
	Branch        string `yaml:"branch"`
	RegexSettings struct {
		UseReg bool   `yaml:",flow"`
		RegStr string `yaml:",flow"`
	}
}

func GetYaml() {
	str := `
path: package.json
behavior: before
template: v{{.Version}}
branch: testbranch1
regexsettings:
  usereg: true
  regstr: 
`
	settings := AtcSettings{}
	yaml.Unmarshal([]byte(str), &settings)
	if !settings.RegexSettings.UseReg {
		log.Println("not use regex")
	}
	fmt.Println(settings)
}
