package main

import (
	"os"
	"testing"
)

func TestFindnpmrc(t *testing.T) {
	npmrc, _ := os.ReadFile("gradle.properties")
	gragle := `## gradle.properties

	# Common Android settings
	android.compileSdkVersion=28
	android.applicationId=com.example
	android.targetSdkVersion=28
	android.minSdkVersion=21
	android.versionCode=2
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
	version=1.0.32
	org.gradle.caching=true
	org.gradle.parallel=true
	org.gradle.caching.debug=false
	org.gradle.configureondemand=false
	org.gradle.daemon.idletimeout= 10800000
	org.gradle.console=auto
	#org.gradle.java.home=(path to JDK home)
	#org.gradle.warning.mode=(all,none,summary)
	#org.gradle.workers.max=(max # of worker processes)
	# org.gradle.priority=(low,normal)`
	if findnpmrc(string(npmrc)) == "" {
		t.Error(`string(npmrc ) = `, string(npmrc))
	}
	if findnpmrc(gragle) == "" {
		t.Error(`graddle err`)
	}
}

func TestNonFindnpmrc(t *testing.T) {
	npmrc := `npm config set init.author.name "Hiro Protagonist"
	npm config set init.author.email "hiro@showcrash.io"
	npm config set init.author.url "http://hiro.snowcrash.io"
	npm config set init.license "MIT"`
	if findnpmrc(npmrc) != "" {
		t.Error(`findnpmrc(npmrc) = ""`)
	}
}
