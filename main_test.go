package main

import (
	"log"
	"regexp"
	"testing"
)

const (
	testStr = `versionName "1.0start"
	plugins {
		id 'com.android.application'
		id 'kotlin-android'
		id 'kotlin-kapt'
		id 'dagger.hilt.android.plugin'
	}
	
	android {
		compileSdk 31
	defaultConfig {   
	 as12-."3 
	}
	
	defaultConfig {
		comment
			applicationId "com.plcoding.cryptocurrencyappyt"
			minSdk 21
			targetSdk 31
			versionCode 1
	//        versionName "1comment"
			versionName "1aftercomment"

	
			testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"
		}
	
		defaultConfig {
			versionCode 1
			versionName "2before"
		}
		productFlavors {
			demo {
				...
				versionName "1.1-demo"
			}
			full {
				...
			}
		}

		defaultConfig {
			1 before 1 after
			applicationId "com.plcoding.cryptocurrencyappyt"
			minSdk 21
			vectorDrawables {
				useSupportLibrary true
			}
			targetSdk 31
			versionCode 1
	//        versionName "3-before and after"
	versionName "1 before 1 after"
			vectorDrawables {
				useSupportLibrary true
			}
	
			testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"
		}

		defaultConfig {
			1 after
			applicationId "com.plcoding.cryptocurrencyappyt"
			minSdk 21
			targetSdk 31
			versionCode 1
	
	versionName "1after"
	//        versionName "1after-comment"
	
			vectorDrawables {
				useSupportLibrary true
				versionName "invector"
			}
	
			testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"
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
)

func TestUnmarshalGradle(t *testing.T) {
	regexStr := `defaultConfig {[^{}]*([^{}]*{[^{}]*}[^{}]*)*[^{}]*\n[\t ]*versionName "(.+)"`
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		t.Error(err)
	}
	res := regex.FindStringSubmatch(testStr)
	if len(res) == 0 {
		t.Error("regex err: ", res)
	}
	for r := range res {
		log.Println(res[r])
	}

}
