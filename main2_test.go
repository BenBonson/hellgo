package main

import (
	"log"
	"testing"
)

func Test_main(t *testing.T) {
	var lan string
	lan_lower := "english"
	if lan_lower == "english" {
		lan = "en"
		log.Println("lan_lower to lan working")
	} else {
		t.Errorf("lan_lower to lan not working")
	}

	if lan == "en" {
		greeting = "hello"
		log.Println("lan to greeting working")
	} else {
		t.Errorf("lan to greeting not working")
	}

	if greeting == "hello" {
		log.Println("English works")
	} else {
		t.Errorf("Expected hello")
	}
}
