package imagefetch

import (
	"math/rand"
	"strings"
	"time"
)

// CharacterCollection represents a collection of character information.
type CharacterCollection struct {
	characters        []CharacterInfo
	reference         map[string]CharacterInfo
	referenceLower    map[string]CharacterInfo
	referenceFriendly map[string]CharacterInfo
}

// Characters returns a slice of all characters in this collection.
func (collection CharacterCollection) Characters() []CharacterInfo {
	return collection.characters
}

// Character returns a CharacterInfo from search string if found.
func (collection CharacterCollection) Character(search string, caseSensitive bool) (character CharacterInfo, ok bool) {
	if caseSensitive {
		character, ok = collection.reference[search]
	} else {
		character, ok = collection.referenceLower[strings.ToLower(search)]
	}
	return
}

// CharacterFriendly returns a CharacterInfo by friendly name.
func (collection CharacterCollection) CharacterFriendly(friendly string) (character CharacterInfo, ok bool) {
	character, ok = collection.referenceFriendly[strings.ToLower(friendly)]
	return
}

// CharacterFullName returns a CharacterInfo by full name.
func (collection CharacterCollection) CharacterFullName(name string) (character CharacterInfo, ok bool) {
	return collection.Character(strings.Replace(name, " ", "_", -1), false)
}

// CharacterRandom returns a random CharacterInfo.
func (collection CharacterCollection) CharacterRandom() (character CharacterInfo, ok bool) {
	switch len(collection.characters) {
	case 0:
		return CharacterInfo{}, false
	case 1:
		return collection.characters[0], true
	}
	rand.Seed(time.Now().UnixNano())
	return collection.characters[rand.Intn(len(collection.characters)-1)], true
}

// CharacterInfo represents information on a character.
type CharacterInfo struct {
	SearchString string
	Color        int
	FriendlyName string
}

// CharacterArt represents a piece of artwork of a character.
type CharacterArt struct {
	ImageURL  string
	SourceURL string
}
