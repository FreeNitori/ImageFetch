package imagefetch

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// FetchPosts fetches a posts of a character.
func FetchPosts(query string, character CharacterInfo) (BoardPosts, error) {
	var (
		response *http.Response
		data     []byte
		result   BoardPosts
	)
	response, err = http.Get(query + character.SearchString)
	if err != nil {
		return BoardPosts{}, err
	}
	defer func() { _ = response.Body.Close() }()
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return BoardPosts{}, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return BoardPosts{}, err
	}
	if result.Count < 1 {
		return BoardPosts{}, ErrNoArtAvailable
	}
	return result, nil
}

// FetchRandomPost fetches a random post.
func FetchRandomPost(posts BoardPosts) (post BoardPost) {
	rand.Seed(time.Now().UnixNano())
	if posts.Count < 100 {
		if posts.Count == 1 {
			post = posts.Posts[0]
		} else {
			post = posts.Posts[rand.Intn(posts.Count-1)]
		}
	} else {
		post = posts.Posts[rand.Intn(99)]
	}
	return
}

// FetchRandomCharacter fetches a random character from a collection.
func FetchRandomCharacter(collection CharacterCollection) (CharacterInfo, error) {
	switch len(collection.characters) {
	case 0:
		return CharacterInfo{}, ErrNoArtAvailable
	case 1:
		return collection.characters[0], nil
	}
	rand.Seed(time.Now().UnixNano())
	return collection.characters[rand.Intn(len(collection.characters)-1)], nil
}

// Fetch fetches a random art of a character.
func Fetch(query string, character CharacterInfo) (CharacterArt, error) {
	var (
		posts BoardPosts
		post  BoardPost
	)
	posts, err = FetchPosts(query, character)
	if err != nil {
		return CharacterArt{}, err
	}
	post = FetchRandomPost(posts)
	return post.Art(), nil
}

// FetchRandom fetches art of a random character in a collection.
func FetchRandom(query string, collection CharacterCollection) (CharacterArt, error) {
	var (
		character CharacterInfo
		posts     BoardPosts
	)
	character, err = FetchRandomCharacter(collection)
	if err != nil {
		return CharacterArt{}, err
	}
	posts, err = FetchPosts(query, character)
	if err != nil {
		return CharacterArt{}, err
	}
	return FetchRandomPost(posts).Art(), nil
}
