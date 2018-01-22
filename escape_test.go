package gourlescape

import "testing"

func TestValidEscapes(t *testing.T) {

	e := Escape{}

	value := e.URL(" Some URL escaP3d text!!  	")

	if value != "some-url-escap3d-text--" {
		t.Fatal("Failed to URL escape test #1")
	}

	value = e.URL("")

	if value != "" {
		t.Fatal("Failed to URL escape empty string")
	}

	value = e.URL("already-escaped-text")

	if value != "already-escaped-text" {
		t.Fatal("Failed to URL escape already escaped text")
	}

	value = e.URL("ALL CAPITAL-LettERS@!")

	if value != "all-capital-letters--" {
		t.Fatal("Failed to escape test #4")
	}
}

func TestDifferentReplacer(t *testing.T) {
	e := Escape{
		Replacer: "_",
	}

	value := e.URL("Test URL with New REPLACER")

	if value != "test_url_with_new_replacer" {
		t.Fatal("Failed to escape underscore replacer")
	}
}

func TestRemoveTail(t *testing.T) {
	e := Escape{
		RemoveTail: true,
	}

	value := e.URL("Test Remove Tail!!$")

	if value != "test-remove-tail" {
		t.Fatal("Failed to remove tail characters.")
	}
}
