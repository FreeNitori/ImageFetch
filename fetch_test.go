package imagefetch

import (
	"testing"
)

func TestTouhouFetchRandom(t *testing.T) {
	art, err := FetchRandom(SafeTouhouQuery, Touhou)
	if err != nil {
		t.Fatalf("Unable to fetch random art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
}

func TestTouhouFetchLookupCaseSensitive(t *testing.T) {
	character, ok := Touhou.Character("Kawashiro_Nitori", true)
	if !ok {
		t.Fatal("Unable to lookup character.")
	}
	art, err := Fetch(SafeTouhouQuery, character)
	if err != nil {
		t.Fatalf("Unable to fetch character art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
	_, ok = Touhou.Character("kawashiro_nitori", true)
	if ok {
		t.Fatal("Fetched image with wrong case.")
	}
}

func TestTouhouFetchLookupCaseInsensitive(t *testing.T) {
	character, ok := Touhou.Character("KAWASHIRO_NITORI", false)
	if !ok {
		t.Fatal("Unable to lookup character.")
	}
	art, err := Fetch(SafeTouhouQuery, character)
	if err != nil {
		t.Fatalf("Unable to fetch character art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
}

func TestTouhouFetchFriendlyName(t *testing.T) {
	character, ok := Touhou.CharacterFriendly("Nitori")
	if !ok {
		t.Fatal("Unable to lookup character.")
	}
	art, err := Fetch(SafeTouhouQuery, character)
	if err != nil {
		t.Fatalf("Unable to fetch character art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
}

func TestTouhouFetchFriendlyNameAlternateCase(t *testing.T) {
	character, ok := Touhou.CharacterFriendly("nitori")
	if !ok {
		t.Fatal("Unable to lookup character.")
	}
	art, err := Fetch(SafeTouhouQuery, character)
	if err != nil {
		t.Fatalf("Unable to fetch character art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
}

func TestTouhouFetchFullName(t *testing.T) {
	character, ok := Touhou.CharacterFullName("Kawashiro Nitori")
	if !ok {
		t.Fatal("Unable to lookup character.")
	}
	art, err := Fetch(SafeTouhouQuery, character)
	if err != nil {
		t.Fatalf("Unable to fetch character art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
}

func TestTouhouFetchFullNameAlternateCase(t *testing.T) {
	character, ok := Touhou.CharacterFullName("kawashiro nitori")
	if !ok {
		t.Fatal("Unable to lookup character.")
	}
	art, err := Fetch(SafeTouhouQuery, character)
	if err != nil {
		t.Fatalf("Unable to fetch character art, %s", err)
	}
	t.Logf("Image URL: %s", art.ImageURL)
	t.Logf("Source URL: %s", art.SourceURL)
}
