package imagefetch

import "strings"

func NewCollection(characters []CharacterInfo) CharacterCollection {
	collection := CharacterCollection{
		characters:        characters,
		reference:         map[string]CharacterInfo{},
		referenceLower:    map[string]CharacterInfo{},
		referenceFriendly: map[string]CharacterInfo{},
	}
	// map references
	for _, character := range collection.characters {
		collection.reference[character.SearchString] = character
		collection.referenceLower[strings.ToLower(character.SearchString)] = character
		collection.referenceFriendly[strings.ToLower(character.FriendlyName)] = character
	}
	return collection
}
